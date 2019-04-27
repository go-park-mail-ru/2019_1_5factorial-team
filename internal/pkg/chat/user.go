package chat

import (
	"context"
	"fmt"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
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

func NewUserID(conn *websocket.Conn, ID string) (*User, error) {
	u, err := user.GetUserById(ID)
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

func NewUser(conn *websocket.Conn, token string) (*User, error) {
	grpcConn, err := grpcAuth.CreateConnection()
	if err != nil {
		log.Error(errors.Wrap(err, "cant connect to auth grpc, NewUser"))
		return nil, nil
	}
	defer grpcConn.Close()

	AuthGRPC := grpcAuth.NewAuthCheckerClient(grpcConn)
	ctx := context.Background()
	uID, err := AuthGRPC.GetIDFromSession(ctx, &grpcAuth.Cookie{Token: token, Expiration: ""})
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetID"))
		return nil, nil
	}
	//ID, err := session.GetId(token)
	//if err != nil {
	//	log.Error(errors.Wrap(err, "cant create user, GetID"))
	//	return nil, nil
	//}

	u, err := user.GetUserById(uID.ID)
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetUserById"))
		return nil, nil
	}

	return &User{
		conn: conn,
		ID:   u.ID.Hex(),
		//Token:      token,
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
			log.Printf("%s, stop listen", u.ID)
			return

		default:
			//log.Printf("player %s ListenMessage default", p.Token)

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
				message.From = u.Nickname
				message.Time = time.Now()
				
				err = message.Insert()
				if err != nil {
					// TODO(): отправить юзеру сообщение, что мессаж не отправился
					u.SendErr(err.Error())
					continue
				}
			case string(MessageTyping):
				message.From = u.Nickname

			}

			
			

			//message.Type = ""
			u.in <- message

			//message := &Message{}
			//err := u.conn.ReadJSON(message)
			//if websocket.IsUnexpectedCloseError(err) || websocket.IsCloseError(err) {
			//	u.ChatPtr.RemoveUser(u)
			//	log.Printf("user %s disconnected", u.ID)
			//
			//	return
			//
			//} else if err != nil {
			//	log.Printf("cannot read json, err = %s", err.Error())
			//	u.SendErr(err.Error())
			//
			//	if e, ok := err.(*net.OpError); ok {
			//		if e.Temporary() || e.Timeout() {
			//			// I don't think these actually happen, but we would want to continue if they did...
			//			continue
			//		} else if e.Err.Error() == "use of closed network connection" { // happens very frequently
			//			// не знаю что тут сделать, выкинуть его из комнаты или шо?
			//			u.stopListen <- struct{}{}
			//			continue
			//		}
			//	}
			//
			//	continue
			//}
			//
			//log.Println(message)
			//if message.Payload == nil {
			//	u.SendErr("empty payload")
			//	continue
			//}
			//
			//if message.Type == MessageNew {
			//	payloadMap := message.Payload.(map[string]interface{})
			//
			//	fmt.Println("CHECK NEW")
			//	um := UserMessage{
			//		From: u.Nickname,
			//		Time: time.Now(),
			//		Text: payloadMap["text"].(string),
			//	}
			//
			//	err := um.Insert()
			//	if err != nil {
			//		// TODO(): отправить юзеру сообщение, что мессаж не отправился
			//		u.SendErr(err.Error())
			//		continue
			//	}
			//	u.in <- &um
			//} else if message.Type == MessageTyping {
			//	log.Println(u.ID, "typing")
			//}

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
	mes, err := GetLastMessages(LastMessagesLimit)
	if err != nil {
		log.Error(errors.Wrapf(err, "user %s cant get last messages on connect", u.Nickname))
	}

	for _, val := range mes {
		err := u.conn.WriteJSON(Message{
			Type:    MessageExist,
			Payload: val,
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
