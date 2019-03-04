package user

import (
	"fmt"
	"sync"
)

type DatabaseUser struct {
	Id int
	Email string
	Nickname string
	Password string
	Score int
	AvatarLink string
}

var once sync.Once
var users map[string]DatabaseUser

func init() {
	once.Do(func() {
		fmt.Println("init users map")
		//users = make([]DatabaseUser, 0)
		users = make(map[string]DatabaseUser)
		users["kek"] = DatabaseUser{
			Id: 1,
			Email: "kek.k.ek",
			Nickname: "kek",
			Password: "password",
			Score: 100500,
			AvatarLink: "./avatars/default.jpg"}

	})
}

func getUsers() map[string]DatabaseUser {
	fmt.Println(users)

	return users
}

func addUser(in DatabaseUser) error {
	// TODO(): add mutex
	users[in.Nickname] = in
}

func PrintUsers() {
	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")
}

