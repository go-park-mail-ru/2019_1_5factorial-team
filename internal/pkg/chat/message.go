package chat

import "time"

type MessageType string

const (
	MessageNew     MessageType = "NEW"
	MessageErr     MessageType = "ERR"
	MessageConnect MessageType = "CONNECTED"
	MessageEnd     MessageType = "END"
)

type UserMessage struct {
	From string    `json:"from"`
	Time time.Time `json:"time"`
	Text string    `json:"text"`
}

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
