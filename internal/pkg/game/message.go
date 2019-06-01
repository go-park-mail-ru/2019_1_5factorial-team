package game

type MessageType string

const (
	MessageState   MessageType = "STATE"
	MessageErr     MessageType = "ERR"
	MessageMove    MessageType = "MOVE"
	MessageConnect MessageType = "CONNECTED"
	MessageEnd     MessageType = "END"
	MessageLink    MessageType = "LINK"
	MessagePause   MessageType = "PAUSE"
	MessageResume  MessageType = "RESUME"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type IncomeMessage struct {
	Type    MessageType `json:"type"`
	Pressed string      `json:"pressed"`
}

//{"type": "MOVE", "pressed": "up"}
