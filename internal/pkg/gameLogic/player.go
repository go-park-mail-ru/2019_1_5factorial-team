package gameLogic

import (
	"context"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
)

type PlayerCharacter struct {
	Token string `json:"-"`
	HP    int    `json:"hp"`
	Score int    `json:"score"`
	Nick  string `json:"nick"`
}

func NewPlayerCharacter(token string, grpcClient grpcAuth.AuthCheckerClient) (PlayerCharacter, error) {
	pc := PlayerCharacter{
		Token: token,
		HP:    5,
		Score: 0,
	}

	ctx := context.Background()
	uID, err := grpcClient.GetIDFromSession(ctx, &grpcAuth.Cookie{Token: token, Expiration: ""})
	if err != nil {
		log.Error(errors.Wrap(err, "cant create user, GetID"))
		return PlayerCharacter{}, nil
	}

	u, err := grpcClient.GetUserByID(ctx, &grpcAuth.User{ID: uID.ID})
	if err != nil {
		return PlayerCharacter{}, errors.Wrap(err, "cant create player character")
	}

	pc.Nick = u.Nickname

	return pc, nil
}
