package game

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Player struct {
	conn *websocket.Conn
	// если нужно хранить всю инфу по пользователю, то хранить User
	ID         string
	Score      int
	room       *Room
	in         chan *IncomeMessage
	out        chan *Message
	unregister chan struct{}
	stopListen chan struct{}
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn:       conn,
		ID:         id,
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
					logrus.Printf("player %s disconnected", p.ID)
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
			p.conn.Close()

		case message := <-p.out:
			p.conn.WriteJSON(message)

		case message := <-p.in:
			logrus.Printf("income: %#v", message)
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
