package user

import (
	"github.com/pkg/errors"
	"math/rand"
)

type User struct {
	Id int
	Email string
	Nickname string
	Password string
	Score int
	AvatarLink string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию
	rid := rand.Int()
	err := addUser(DatabaseUser{
		Id: rid,
		Email: email,
		Nickname: nickname,
		Password: password,
	})
	if err != nil {
		err = errors.Wrap(err, "Cannot create user")
		return User{}, err
	}

	u := User{
		Id: rid,
		Email: email,
		Nickname: nickname,
		Password: password,
	}

	return u, nil
}
