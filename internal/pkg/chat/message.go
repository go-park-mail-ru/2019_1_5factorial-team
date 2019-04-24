package chat

import (
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

type UserMessage struct {
	ID   bson.ObjectId `json:"id"     bson:"_id,omitempty"`
	From string        `json:"from"   bson:"from"`
	Time time.Time     `json:"time"   bson:"time"`
	Text string        `json:"text"   bson:"text"`
}

type ErrMessage struct {
	Error string `json:"error"`
}

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
