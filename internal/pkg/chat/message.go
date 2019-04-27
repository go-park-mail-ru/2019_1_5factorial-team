package chat

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MessageType string

const (
	MessageNew    MessageType = "NEW"
	MessageDelete MessageType = "DELETE"
	MessageEdit   MessageType = "EDIT"
	MessageReply  MessageType = "REPLY"
	MessageTyping MessageType = "TYPING"

	MessageErr     MessageType = "ERR"
	MessageConnect MessageType = "CONNECTED"
	MessageExist   MessageType = "EXIST"
)

var validUserInput = map[string]MessageType{
	string(MessageNew):    MessageNew,
	string(MessageDelete): MessageDelete,
	string(MessageEdit):   MessageEdit,
	string(MessageReply):  MessageReply,
	string(MessageTyping): MessageTyping,
}

// TODO(): type должен быть в инкаме, а в рассылках не указываться
type UserMessage struct {
	Type string        `json:"type,omitempty"   bson:"-"`
	ID   bson.ObjectId `json:"id,omitempty"     bson:"_id,omitempty"`
	From string        `json:"from,omitempty"   bson:"from"`
	Time time.Time     `json:"time,omitempty"   bson:"time"`
	Text string        `json:"text,omitempty"   bson:"text"`

	DeleteID string `json:"delete_id,omitempty" bson:"-"`
}

func (um *UserMessage) SetMessageNew() error {
	um.DeleteID = ""

	if um.Text == "" || um.Text == " " {
		return errors.New("empty NEW message text")
	}

	return nil
}

func (um *UserMessage) SetMessageTyping() {
	um.ID = ""
	um.Text = ""
	um.DeleteID = ""
}

func (um *UserMessage) SetMessageDelete() error {
	um.ID = ""
	um.Text = ""

	if um.DeleteID == "" || um.DeleteID == " " {
		return errors.New("delete id is empty")
	}

	return nil
}

func (um *UserMessage) Validate() error {
	if _, ok := validUserInput[um.Type]; !ok {
		return errors.New("not valid message type")
	}

	//if um.Text == "" || um.Text == " " {
	//	return errors.New("empty message payload")
	//}

	return nil
}

type ErrMessage struct {
	Error string `json:"error"`
}

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
