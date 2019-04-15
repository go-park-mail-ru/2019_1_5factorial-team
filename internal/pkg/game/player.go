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

func (p *Player) Listen() {
	go func() {
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
	}()

	for {
		select {
		case <-p.unregister:
			_ = p.conn.Close()

		case message := <-p.out:
			_ = p.conn.WriteJSON(message)

		case message := <-p.in:
			logrus.Printf("from player = %s, income: %#v", p.Token, message)

			if message.Type != "MOVE" {
				logrus.Println("not move")
				_ = p.conn.WriteJSON(Message{
					Type: "ERR",
					Payload: "not valid input",
				})

				continue
			}

			button, err := gameLogic.MatchSymbol(message.Pressed)
			if err != nil {
				_ = p.conn.WriteJSON(Message{
					Type: "ERR",
					Payload: "not valid input",
				})

				continue
			}
			logrus.Println(button)
		}
	}
}

func (p *Player) SendState(state *RoomState) {
	p.out <- &Message{"STATE", state}
}

func (p *Player) SendMessage(message *Message) {
	p.out <- message
}

func (p *Player) CloseConn() {
	p.unregister <- struct{}{}
	p.stopListen <- struct{}{}
	//_ = p.conn.Close()
}
