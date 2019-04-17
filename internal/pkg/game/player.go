package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gameLogic"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Player struct {
	conn *websocket.Conn
	// если нужно хранить всю инфу по пользователю, то хранить User
	Token      string
	Score      int
	room       *Room
	in         chan *IncomeMessage
	out        chan *Message
	unregister chan struct{}
	stopListen chan struct{}
}

func NewPlayer(conn *websocket.Conn, token string) *Player {
	return &Player{
		conn:       conn,
		Token:      token,
		Score:      0,
		in:         make(chan *IncomeMessage),
		out:        make(chan *Message),
		unregister: make(chan struct{}),
		stopListen: make(chan struct{}),
	}
}

func (p *Player) ListenMessages() {
	for {
		select {
		case <-p.stopListen:
			return

		default:
			message := &IncomeMessage{}
			err := p.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) {
				p.room.RemovePlayer(p)
				logrus.Printf("player %s disconnected", p.Token)
				return
			}
			if err != nil {
				logrus.Printf("cannot read json")
				continue
			}

			p.in <- message
		}
	}
}

func (p *Player) Listen() {
	go p.ListenMessages()

	for {
		select {
		case <-p.unregister:
			err := p.conn.Close()
			if err != nil {
				logrus.Error("p.Listen cant close connection", err.Error())
			}

		case message := <-p.out:
			err := p.conn.WriteJSON(message)
			if err != nil {
				logrus.Error("p.Listen cant send message", err.Error())
				p.CloseConn()
			}

		case message := <-p.in:
			logrus.Printf("from player = %s, income: %#v", p.Token, message)

			if message.Type != "MOVE" {
				logrus.Println("not move")
				err := p.conn.WriteJSON(Message{
					Type:    MessageErr,
					Payload: "not valid input",
				})

				if err != nil {
					logrus.Error("p.Listen cant send message", err.Error())
				}
				p.CloseConn()

				continue
			}

			button, err := gameLogic.MatchSymbol(message.Pressed)
			if err != nil {
				err = p.conn.WriteJSON(Message{
					Type:    MessageErr,
					Payload: "not valid input",
				})

				if err != nil {
					logrus.Error("p.Listen cant send message", err.Error())
				}
				p.CloseConn()

				continue
			}
			logrus.Println(button)
			p.room.playerInput <- &button
		}
	}
}

func (p *Player) SendState(state *RoomState) {
	p.out <- &Message{
		Type:    MessageState,
		Payload: state,
	}
}

func (p *Player) SendMessage(message *Message) {
	p.out <- message
}

func (p *Player) CloseConn() {
	p.unregister <- struct{}{}
	p.stopListen <- struct{}{}
	//_ = p.conn.Close()
}
