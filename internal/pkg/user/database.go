package user

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"sync/atomic"
)

type DatabaseUser struct {
	Id int64
	Email string
	Nickname string
	HashPassword string
	Score int64
	AvatarLink string
}

var once sync.Once
var users map[string]DatabaseUser
var mu *sync.Mutex
var currentId int64

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
	defer mu.Unlock()
	mu.Lock()
	
	fmt.Println(users)

	return users
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
	defer mu.Unlock()
	mu.Lock()

	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")
}

func getNextId() int64 {
	atomic.AddInt64(&currentId, 1)

	return currentId
}