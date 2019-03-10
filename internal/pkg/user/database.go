package user

import (
	"fmt"
	"github.com/pkg/errors"
	"sort"
	"sync"
)

type DatabaseUser struct {
	Id           int64
	Email        string
	Nickname     string
	HashPassword string
	Score        int
	AvatarLink   string
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
		users["kekkekkek"] = DatabaseUser{
			Id:           0,
			Email:        "kek.k.ek",
			Nickname:     "kekkekkek",
			HashPassword: hash,
			Score:        100500,
			AvatarLink:   "./avatars/default.jpg",
		}

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

func findUserById(id int64) (DatabaseUser, error) {
	defer mu.Unlock()

	mu.Lock()
	for _, val := range users {
		if val.Id == id {
			return val, nil
		}
	}

	return DatabaseUser{}, errors.New("user with this id not found")
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

func getNextId() int64 {
	defer mu.Unlock()
	mu.Lock()
	currentId++

	return currentId
}

func updateDBUser(user DatabaseUser) error {
	defer mu.Unlock()

	mu.Lock()
	if _, ok := users[user.Nickname]; !ok {
		return errors.New("cannot find user")
	}

	users[user.Nickname] = user
	return nil
}

type ByNameScore []DatabaseUser

func (a ByNameScore) Len() int			 { return len(a) }
func (a ByNameScore) Less(i, j int) bool { return a[i].Score < a[j].Score }
func (a ByNameScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getScores() []DatabaseUser {
	mu.Lock()
	result := make([]DatabaseUser, 0, 1)
	for _, val := range users {
		result = append(result, val)
	}

	sort.Sort(ByNameScore(result))

	mu.Unlock()
	return result
}

func getUsersCount() int {
	return len(users)
}
