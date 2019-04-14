package gameLogic

type PlayerCharacter struct {
	ID     string `json:"id"`
	Sprite string `json:"sprite"`
	X      int    `json:"x"`
	HP     int    `json:"hp"`
}

func NewPlayerCharacter(id string) PlayerCharacter {
	return PlayerCharacter{
		ID:     id,
		Sprite: "default",
		X:      0,
		HP:     100,
	}
}
