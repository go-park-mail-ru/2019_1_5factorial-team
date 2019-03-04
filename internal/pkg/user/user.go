package user

import (
	"github.com/pkg/errors"
	"math/rand"
)

type User struct {
	Id int
	Email string
	Nickname string
	HashPassword string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию на повторение ника и почты

	rid := rand.Int()
	hashPassword, err := getPasswordHash(password)
	if err != nil {
		return User{}, errors.Wrap(err, "Hasher password error")
	}

	err = addUser(DatabaseUser{
		Id: rid,
		Email: email,
		Nickname: nickname,
		HashPassword: hashPassword,
	})
	if err != nil {
		err = errors.Wrap(err, "Cannot create user")
		return User{}, err
	}

	u := User{
		Id: rid,
		Email: email,
		Nickname: nickname,
		HashPassword: hashPassword,
	}

	return u, nil
}
