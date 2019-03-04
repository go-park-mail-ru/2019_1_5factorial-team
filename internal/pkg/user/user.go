package user

import (
	"github.com/pkg/errors"
)

type User struct {
	Id int
	Email string
	Nickname string
	HashPassword string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию на повторение ника и почты

	rid := getNextId()
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

func IdentifyUser(login string, password string) (User, error) {
	u, err := getUser(login)
	if err != nil {
		return User{}, errors.Wrap(err, "Can't find user")
	}

	err = comparePassword(password, u.HashPassword)
	if err != nil {
		return User{}, errors.Wrap(err, "Wrong password")
	}

	return User{
		Id: u.Id,
		Email: u.Email,
		Nickname: u.Nickname,
	}, nil
}
