package game

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type IncomeMessage struct {
	Type string `json:"type"`
	//Player  string `json:"player"`
	Pressed string `json:"pressed"`
}

//{"type": "MOVE", "pressed": "up"}
