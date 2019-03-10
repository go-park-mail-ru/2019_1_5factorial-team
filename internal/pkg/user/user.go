package user

import (
	"github.com/pkg/errors"
)

type User struct {
	Id           int64
	Email        string
	Nickname     string
	HashPassword string
	Score        int
	AvatarLink   string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию на повторение ника и почты

	rid := getNextId()
	hashPassword, err := getPasswordHash(password)
	if err != nil {
		return User{}, errors.Wrap(err, "Hasher password error")
	}

	err = addUser(DatabaseUser{
		Id:           rid,
		Email:        email,
		Nickname:     nickname,
		HashPassword: hashPassword,
	})
	if err != nil {
		err = errors.Wrap(err, "Cannot create user")
		return User{}, err
	}

	u := User{
		Id:           rid,
		Email:        email,
		Nickname:     nickname,
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
		Id:       u.Id,
		Email:    u.Email,
		Nickname: u.Nickname,
	}, nil
}

func GetUserById(id int64) (User, error) {
	u, err := findUserById(id)
	if err != nil {
		return User{}, errors.Wrap(err, "can't find user")
	}

	return User{
		Id:         u.Id,
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      u.Score,
		AvatarLink: u.AvatarLink,
	}, nil
}

func UpdateUser(id int64, oldPassword string, newPassword string) error {
	u, err := findUserById(id)
	if err != nil {
		return errors.Wrap(err, "update user error")
	}

	// не будет вложенных ифов, пока не решим менять логин
	if newPassword == "" {
		return errors.New("nothing to update")
	}

	err = validateChangingPasswords(oldPassword, newPassword, u.HashPassword)
	if err != nil {
		return errors.Wrap(err, "validate passwords error")
	}

	newHashPassword, err := getPasswordHash(newPassword)
	if err != nil {
		return errors.Wrap(err, "some password error")
	}

	u.HashPassword = newHashPassword
	err = updateDBUser(u)
	if err != nil {
		return errors.Wrap(err, "cant update avatar")
	}

	return nil
}

func validateChangingPasswords(oldPassword string, newPassword string, currentHashPassword string) error {
	if oldPassword == "" {
		return errors.New("please input old password")
	}

	err := comparePassword(newPassword, currentHashPassword)
	if err == nil {
		return errors.New("old and new password are same")
	}

	err = comparePassword(oldPassword, currentHashPassword)
	if err != nil {
		return errors.Wrap(err, "invalid old password")
	}

	return nil
}
