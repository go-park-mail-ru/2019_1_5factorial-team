package donation

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"sync"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/random"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

// TODO: user listener and writer

type User struct {
	id               string
	conn             *websocket.Conn
	out              chan *string
	deletionCallback func()
}

func NewUser(conn *websocket.Conn) (User, error) {
	uid, err := random.RandomUUID()
	if err != nil {
		return User{}, errors.Wrap(err, "failed to create a user")
	}
	return User{
		id:   uid,
		conn: conn,
		out:  make(chan *string, 10),
	}, nil
}

func (u *User) Listen() {
	for {
		select {
		case message := <-u.out:
			err := u.conn.WriteJSON(Message{
				Type: "NEW",
				Msg:  *message,
			})

			if err != nil {
				log.Error("u.Listen cant send message, closing connection", err.Error())

				err := u.conn.Close()
				log.Printf("close connection on user %v", u)
				if err != nil {
					log.Error("u.Listen cant close connection", err.Error())
				}

				u.deletionCallback()

				return
			}
		}
	}
}

// thread - safe users map
type UsersMap struct {
	mu    *sync.Mutex
	users map[string]User
}

func NewUsersMap() UsersMap {
	return UsersMap{
		mu:    &sync.Mutex{},
		users: make(map[string]User),
	}
}

func (users UsersMap) Add(u User) {
	users.mu.Lock()
	u.deletionCallback = func() {
		users.Delete(u.id)
	}
	users.users[u.id] = u
	users.mu.Unlock()
}

func (users UsersMap) Delete(id string) {
	users.mu.Lock()
	delete(users.users, id)
	users.mu.Unlock()
}

func (users UsersMap) SendMessages(msg string) {
	users.mu.Lock()
	for _, u := range users.users {
		u.out <- &msg
	}
	users.mu.Unlock()
}
