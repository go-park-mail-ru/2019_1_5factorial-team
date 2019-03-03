package user

import "math/rand"

type User struct {
	Id int
	Email string
	Nickname string
	Password string
	Score int
	AvatarType string
	AvatarLink string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию
	rid := rand.Int()
	addUser(DatabaseUser{
		id: rid,
		Email: email,
		Nickname: nickname,
		Password: password,
	})

	u := User{
		Id: rid,
		Email: email,
		Nickname: nickname,
		Password: password,
	}

	return u, nil
}
