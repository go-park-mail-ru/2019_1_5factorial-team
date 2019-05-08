package chat

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"sync"
	"time"
)

var InstanceChat *Chat

func Start() {
	//чат крутится как отдельная сущность всегда
	InstanceChat = NewChat(config.Get().ChatConfig.MaxUsers)
	go panicWorker.PanicWorker(InstanceChat.Start)
}

type Chat struct {
	mu           *sync.Mutex
	UserChan     chan *User
	Users        map[string]*User
	MaxUsers     int
	register     chan *User
	unregister   chan *User
	ticker       *time.Ticker
	messagesChan chan UserMessage
	messages     []Message
}

func NewChat(maxUsers int) *Chat {
	return &Chat{
		mu:           &sync.Mutex{},
		UserChan:     make(chan *User, 10),
		Users:        make(map[string]*User),
		MaxUsers:     maxUsers,
		register:     make(chan *User, 10),
		unregister:   make(chan *User, 10),
		ticker:       time.NewTicker(1 * time.Second),
		messagesChan: make(chan UserMessage, 10),
		messages:     make([]Message, 0, 100),
	}
}

func (c *Chat) Start() {
	log.Printf("chat started, max users = %d", c.MaxUsers)

	for {
		select {
		case unregisterUser := <-c.unregister:
			log.Printf("unregister user %s, Chat.Start()", unregisterUser.ID)
			delete(c.Users, unregisterUser.Nickname)

		case registerUser := <-c.register:
			log.Printf("register user %s, Chat.Start()", registerUser.ID)
			registerUser.out <- &Message{
				Type:    MessageConnect,
				Payload: nil,
			}

			registerUser.ChatPtr = c
			registerUser.SendLastMessages()
			c.Users[registerUser.Nickname] = registerUser

		case <-c.ticker.C:
			c.BroadcastSendMessages()

		case message := <-c.messagesChan:
			// костыли пезда (но работает)
			// а локалке Payload нормально кастится в map[string]interface{}, а на тачке хуй (хз поч)
			switch message.Type {
			case string(MessageNew):
				message.Type = ""
				c.messages = append(c.messages, Message{
					Type:    MessageNew,
					Payload: message,
				})
			case string(MessageTyping):
				message.Type = ""
				c.messages = append(c.messages, Message{
					Type:    MessageTyping,
					Payload: message,
				})
			case string(MessageDelete):
				message.Type = ""
				c.messages = append(c.messages, Message{
					Type:    MessageDelete,
					Payload: message,
				})
			}

		}
	}
}

func (c *Chat) AddUser(user *User) {
	if len(c.Users) >= c.MaxUsers {
		log.Printf("user %s cant add, too many connects, connects = %d, max = %d", user.ID, len(c.Users), c.MaxUsers)

		user.out <- &Message{
			Type: MessageErr,
			Payload: ErrMessage{
				Error: "sorry, too many clients",
			},
		}
		user.ChatPtr = c
		user.CloseConn()

		return
	}
	c.register <- user
}

func (c *Chat) RemoveUser(user *User) {
	c.unregister <- user
}

func (c *Chat) BroadcastSendMessages() {
	if len(c.messages) == 0 {
		return
	}

	for _, mes := range c.messages {
		for _, user := range c.Users {
			// send mess to user
			log.Printf("send mes = %v to user = %s", mes, user.Nickname)
			user.out <- &mes
		}
	}

	c.messages = make([]Message, 0, 100)
}
