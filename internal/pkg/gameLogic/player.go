package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
)

type PlayerCharacter struct {
	Token string `json:"id"`
	//Sprite string `json:"sprite"`
	//X      int    `json:"x"`
	HP    int    `json:"hp"`
	Score int    `json:"score"`
	Nick  string `json:"nick"`
}

func NewPlayerCharacter(token string) (PlayerCharacter, error) {
	pc := PlayerCharacter{
		Token: token,
		//Sprite: "default",
		//X:      0,
		HP:    100,
		Score: 0,
	}

	id, err := session.GetId(token)
	if err != nil {
		return PlayerCharacter{}, errors.Wrap(err, "cant create player character")
	}

	u, err := user.GetUserById(id)
	if err != nil {
		return PlayerCharacter{}, errors.Wrap(err, "cant create player character")
	}

	pc.Nick = u.Nickname

	return pc, nil
}
