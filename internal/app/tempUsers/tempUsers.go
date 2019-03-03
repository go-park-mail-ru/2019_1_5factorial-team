package tempUsers

import (
	"fmt"
	"sync"
)

type User struct {
	Email string
	Nickname string
	Password string
	Score int
	AvatarType string
	AvatarLink string
}

var users []User
var once sync.Once

func init() {
	once.Do(func() {
		fmt.Println("once do")
		users = make([]User, 0)
		users = append(users, User{
			Email: "kek.k.ek",
			Nickname: "kek",
			Password: "password",
			Score: 100500,
			AvatarType: "jpg",
			AvatarLink: "./avatars/default.jpg"})
	})
}

func GetUsers() []User {
	fmt.Println(users)

	return users
}

func AddUser(in User) {
	users = append(users, in)
}

func PrintUsers() {
	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")
}