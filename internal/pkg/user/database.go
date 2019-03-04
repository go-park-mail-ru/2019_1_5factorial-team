package user

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type DatabaseUser struct {
	Id int
	Email string
	Nickname string
	HashPassword string
	Score int
	AvatarLink string
}

var once sync.Once
var users map[string]DatabaseUser
var mu *sync.Mutex
var currentId int

func init() {
	once.Do(func() {
		fmt.Println("init users map")

		users = make(map[string]DatabaseUser)

		hash, _ := getPasswordHash("password")
		users["kek"] = DatabaseUser{
			Id: 0,
			Email: "kek.k.ek",
			Nickname: "kek",
			HashPassword: hash,
			Score: 100500,
			AvatarLink: "./avatars/default.jpg"}

		mu = &sync.Mutex{}
		currentId = 0
	})
}

func getUsers() map[string]DatabaseUser {
	fmt.Println(users)

	return users
}

func getUser(login string) (DatabaseUser, error) {
	defer mu.Unlock()

	mu.Lock()
	if _, ok := users[login]; !ok {
		return DatabaseUser{}, errors.New("Invalid login")
	} else {
		return users[login], nil
	}
}

func addUser(in DatabaseUser) error {
	defer mu.Unlock()
	mu.Lock()

	if _, ok := users[in.Nickname]; ok {
		return errors.New("User with this nickname already exist")
	}

	users[in.Nickname] = in

	return nil
}

func PrintUsers() {
	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")
}

func getNextId() int {
	defer mu.Unlock()
	mu.Lock()
	currentId++

	return currentId
}