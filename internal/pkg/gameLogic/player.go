package gameLogic

type PlayerCharacter struct {
	Token string `json:"id"`
	//Sprite string `json:"sprite"`
	//X      int    `json:"x"`
	HP    int `json:"hp"`
	Score int `json:"score"`
}

func NewPlayerCharacter(token string) PlayerCharacter {
	return PlayerCharacter{
		Token: token,
		//Sprite: "default",
		//X:      0,
		HP:    100,
		Score: 0,
	}
}
