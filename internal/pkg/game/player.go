package game

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Player struct {
	conn *websocket.Conn
	// если нужно хранить всю инфу по пользователю, то хранить User
	ID   string
	room *Room
	in   chan *IncomeMessage
	out  chan *Message
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn: conn,
		ID:   id,
		in:   make(chan *IncomeMessage),
		out:  make(chan *Message),
	}
}

func (p *Player) Listen() {
	go func() {
		for {
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
	}()

	for {
		select {
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