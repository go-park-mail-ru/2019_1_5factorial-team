package chat

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net"
	"time"
)

type User struct {
	conn       *websocket.Conn
	ID         string
	Token      string
	Nickname   string
	Avatar     string
	ChatPtr    *Chat
	in         chan *UserMessage
	out        chan *Message
	unregister chan struct{}
	stopListen chan struct{}
}

func NewUser(conn *websocket.Conn, token string) (*User, error) {
	ID, err := session.GetId(token)
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetID"))
		return nil, nil
	}

	u, err := user.GetUserById(ID)
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetUserById"))
		return nil, nil
	}

	return &User{
		conn:       conn,
		ID:         u.ID.Hex(),
		Token:      token,
		Nickname:   u.Nickname,
		Avatar:     u.AvatarLink,
		in:         make(chan *UserMessage),
		out:        make(chan *Message),
		unregister: make(chan struct{}),
		stopListen: make(chan struct{}),
	}, nil
}

func (u *User) ListenIncome() {
	for {
		select {
		case <-u.stopListen:
			log.Println("len stopListen", len(u.stopListen))
			log.Printf("%s, stop listen", u.Token)
			return

		default:
			//log.Printf("player %s ListenMessage default", p.Token)

			message := &UserMessage{}
			err := u.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) || websocket.IsCloseError(err) {
				u.ChatPtr.RemoveUser(u)
				log.Printf("player %s disconnected", u.Token)

				return

			} else if err != nil {
				log.Printf("cannot read json, err = %s", err.Error())

				if e, ok := err.(*net.OpError); ok {
					if e.Temporary() || e.Timeout() {
						// I don't think these actually happen, but we would want to continue if they did...
						continue
					} else if e.Err.Error() == "use of closed network connection" { // happens very frequently
						// не знаю что тут сделать, выкинуть его из комнаты или шо?
						u.stopListen <- struct{}{}
						continue
					}
				}

				continue
			}

			if message.Text == "" || message.Text == " " {
				continue
			}

			message.From = u.Nickname
			message.Time = time.Now()
			u.in <- message
		}
	}
}

func (u *User) Listen() {
	go panicWorker.PanicWorker(u.ListenIncome)

	for {
		select {
		case <-u.unregister:
			err := u.conn.Close()
			log.Printf("close connection on user %#v", u)
			if err != nil {
				log.Error("p.Listen cant close connection", err.Error())
			}

			return

		case message := <-u.out:
			err := u.conn.WriteJSON(message)
			if err != nil {
				log.Error("u.Listen cant send message ", err.Error())

				u.CloseConn()
				//return
			}

		case message := <-u.in:
			log.Printf("from player = %s, income: %#v", u, message)

			u.ChatPtr.messagesChan <- *message
		}
	}
}

func (u *User) CloseConn() {
	u.unregister <- struct{}{}
	//p.stopListen <- struct{}{}
}
