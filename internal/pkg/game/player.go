package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gameLogic"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/websocket"
	"net"
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
			log.Println("len stopListen", len(p.stopListen))
			log.Printf("%s, stop listen", p.Token)
			return

		default:
			//log.Printf("player %s ListenMessage default", p.Token)

			message := &IncomeMessage{}
			err := p.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) || websocket.IsCloseError(err) {
				p.room.RemovePlayer(p)
				log.Printf("player %s disconnected", p.Token)

				return

			} else if err != nil {
				log.Printf("cannot read json, err = %s", err.Error())

				if e, ok := err.(*net.OpError); ok {
					if e.Temporary() || e.Timeout() {
						// I don't think these actually happen, but we would want to continue if they did...
						continue
					} else if e.Err.Error() == "use of closed network connection" { // happens very frequently
						// не знаю что тут сделать, выкинуть его из комнаты или шо?
						p.stopListen <- struct{}{}
						continue
					}
				}

				continue
			}

			p.in <- message
		}
	}
}

func (p *Player) Listen() {
	go PanicWorker(p.ListenMessages)

	for {
		select {
		case <-p.unregister:
			err := p.conn.Close()
			log.Printf("close connection on player %s", p.Token)
			if err != nil {
				log.Error("p.Listen cant close connection", err.Error())
			}

			return

		case message := <-p.out:
			err := p.conn.WriteJSON(message)
			if err != nil {
				log.Error("p.Listen cant send message ", err.Error())

				p.CloseConn()
				//return
			}

		case message := <-p.in:
			log.Printf("from player = %s, income: %#v", p.Token, message)

			if message.Type != "MOVE" {
				log.Println("not valid user input")

				err := p.conn.WriteJSON(Message{
					Type:    MessageErr,
					Payload: "not valid input",
				})
				if err != nil {
					log.Error("p.Listen cant send message before match symbol", err.Error())

					p.CloseConn()
				}

				continue
			}

			button, err := gameLogic.MatchSymbol(message.Pressed)
			if err != nil {
				err = p.conn.WriteJSON(Message{
					Type:    MessageErr,
					Payload: "not valid input",
				})
				if err != nil {
					log.Error("p.Listen cant send message in match symbol", err.Error())

					p.CloseConn()
				}

				continue
			}
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
	//p.stopListen <- struct{}{}
}
