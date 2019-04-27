package chat

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MessageType string

const (
	MessageNew     MessageType = "NEW"
	MessageErr     MessageType = "ERR"
	MessageConnect MessageType = "CONNECTED"
	MessageDelete  MessageType = "DELETE"
	MessageExist   MessageType = "EXIST"
)

var validUserInput = map[string]MessageType{
	string(MessageNew):    MessageNew,
	string(MessageDelete): MessageDelete,
}

// TODO(): type должен быть в инкаме, а в рассылках не указываться
type UserMessage struct {
	Type string        `json:"type,omitempty"   bson:"-"`
	ID   bson.ObjectId `json:"id"     bson:"_id,omitempty"`
	From string        `json:"from"   bson:"from"`
	Time time.Time     `json:"time"   bson:"time"`
	Text string        `json:"text"   bson:"text"`
}

func (um *UserMessage) Validate() error {
	if _, ok := validUserInput[um.Type]; !ok {
		return errors.New("not valid message type")
	}

	if um.Text == "" || um.Text == " " {
		return errors.New("empty message payload")
	}

	return nil
}

type ErrMessage struct {
	Error string `json:"error"`
}

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
