package user

import (
	"fmt"
	"sync"
)

type DatabaseUser struct {
	id int
	Email string
	Nickname string
	Password string
	Score int
	AvatarType string
	AvatarLink string
}

var users []DatabaseUser
var once sync.Once

func init() {
	once.Do(func() {
		fmt.Println("once do")
		users = make([]DatabaseUser, 0)
		users = append(users, DatabaseUser{
			Email: "kek.k.ek",
			Nickname: "kek",
			Password: "password",
			Score: 100500,
			AvatarType: "jpg",
			AvatarLink: "./avatars/default.jpg"})
	})
}

func getUsers() []DatabaseUser {
	fmt.Println(users)

	return users
}

func addUser(in DatabaseUser) {
	users = append(users, in)
}

func PrintUsers() {
	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")
}

