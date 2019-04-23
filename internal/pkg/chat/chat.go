package chat

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"sync"
	"time"
)

var InstanceChat *Chat

func init() {
	// игра крутится как отдельная сущность всегда
	InstanceChat = NewChat(20)
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
			delete(c.Users, unregisterUser.Nickname)

		case registerUser := <-c.register:
			registerUser.ChatPtr = c
			registerUser.SendLastMessages()
			c.Users[registerUser.Nickname] = registerUser

		case <-c.ticker.C:
			c.BroadcastSendMessages()

		case message := <-c.messagesChan:
			c.messages = append(c.messages, Message{
				Type:    MessageNew,
				Payload: message,
			})
		}
	}
}

func (c *Chat) AddUser(user *User) {
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
