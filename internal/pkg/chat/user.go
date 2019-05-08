package chat

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/gorilla/websocket"
	"github.com/icrowley/fake"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"net"
	"strings"
	"time"
)

const LastMessagesLimit = 50

type User struct {
	conn *websocket.Conn
	ID   string
	//Token      string
	Nickname   string
	Avatar     string
	ChatPtr    *Chat
	in         chan *UserMessage
	out        chan *Message
	unregister chan struct{}
	stopListen chan struct{}
}

func NewUserID(conn *websocket.Conn, ID string, GRPC grpcAuth.AuthCheckerClient) (*User, error) {
	ctx := context.Background()
	u, err := GRPC.GetUserByID(ctx, &grpcAuth.User{ID: ID})
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetUserById"))

		return nil, errors.Wrap(err, "cant create user, GetUserById")
	}

	return &User{
		conn:       conn,
		ID:         ID,
		Nickname:   u.Nickname,
		Avatar:     u.AvatarLink,
		in:         make(chan *UserMessage),
		out:        make(chan *Message),
		unregister: make(chan struct{}),
		stopListen: make(chan struct{}),
	}, nil
}

func NewUserFake(conn *websocket.Conn) (*User, error) {
	ID := bson.NewObjectId().Hex()
	FakeNick := getFakeNick()
	FakeAvatar := ""

	return &User{
		conn:       conn,
		ID:         ID,
		Nickname:   FakeNick,
		Avatar:     FakeAvatar,
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
			log.Printf("%s, stop listen", u.ID)
			return

		default:

			message := &UserMessage{}
			err := u.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) || websocket.IsCloseError(err) {
				u.ChatPtr.RemoveUser(u)
				log.Printf("user %s disconnected", u.ID)

				return

			} else if err != nil {
				log.Printf("cannot read json, err = %s", err.Error())
				u.SendErr(err.Error())

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

			if err = message.Validate(); err != nil {
				u.SendErr(err.Error())
				continue
			}

			switch message.Type {
			case string(MessageNew):
				err := message.SetMessageNew()
				if err != nil {
					log.Println("cant SetMessageNew()", err)
					u.SendErr(err.Error())
					continue
				}

				message.From = u.Nickname
				message.Time = time.Now()

				err = message.Insert()
				if err != nil {
					log.Println("cant insert new msg", err)
					u.SendErr(err.Error())
					continue
				}

			case string(MessageTyping):
				message.SetMessageTyping()

				message.From = u.Nickname

			case string(MessageDelete):
				err := message.SetMessageDelete()
				if err != nil {
					log.Println("cant SetMessageDelete()", err)
					u.SendErr(err.Error())
					continue
				}

				message.From = u.Nickname
				err = message.Delete()
				if err != nil {
					log.Println("cant delete msg", err)
					u.SendErr(err.Error())
					continue
				}
			}

			u.in <- message
		}
	}
}

func (u *User) Listen() {
	go panicWorker.PanicWorker(u.ListenIncome)

	for {
		select {
		case <-u.unregister:
			u.ChatPtr.RemoveUser(u)
			err := u.conn.Close()
			log.Printf("close connection on user %v", u)
			if err != nil {
				log.Error("u.Listen cant close connection", err.Error())
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
			log.Printf("from user = %s, income: %v", u.Nickname, message)

			u.ChatPtr.messagesChan <- *message
		}
	}
}

func (u *User) CloseConn() {
	u.unregister <- struct{}{}
	//p.stopListen <- struct{}{}
}

func (u *User) SendErr(error string) {
	u.out <- &Message{
		Type: MessageErr,
		Payload: ErrMessage{
			Error: error,
		},
	}
}

func (u *User) SendLastMessages() {
	mes, err := GetLastMessages(config.Get().ChatConfig.LastMessagesLimit)
	if err != nil {
		log.Error(errors.Wrapf(err, "user %s cant get last messages on connect", u.Nickname))
	}

	for i := len(mes) - 1; i > 0; i-- {
		err := u.conn.WriteJSON(Message{
			Type:    MessageExist,
			Payload: mes[i],
		})
		if err != nil {
			log.Error("u.Listen cant send message ", err.Error())

			u.CloseConn()
			//return
		}
	}
}

func getFakeNick() string {
	color := fake.Color()
	jobTitle := fake.JobTitle()
	jobTitle = strings.Replace(jobTitle, " ", "_", -1)
	resultFakeName := fmt.Sprintf("%s_%s", color, jobTitle)
	return resultFakeName
}
