package user

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/manveru/faker"
	"github.com/pkg/errors"
)

const DefaultAvatarLink = "../../../img/default.jpg"

var once sync.Once
var mu *sync.Mutex
var users map[string]User
var currentId int64

func init() {
	once.Do(func() {
		fmt.Println("init users map")

		fake, _ := faker.New("en")
		fake.Rand = rand.New(rand.NewSource(42))

		users = make(map[string]User)

		hash, _ := getPasswordHash("password")
		users["kekkekkek"] = User{
			Id:           0,
			Email:        "kek.k.ek",
			Nickname:     "kekkekkek",
			HashPassword: hash,
			Score:        100500,
			AvatarLink:   DefaultAvatarLink,
		}
		mu = &sync.Mutex{}
		currentId = 0

		var id int64
		// TODO(smet1): generate fake accs in func
		for i := 0; i < 20; i++ {
			id = getNextId()
			nick := fake.FirstName()
			hash, _ := getPasswordHash(nick)

			fmt.Println("id:", id, ", Nick:", nick, ", Password:", nick)

			users[nick] = User{
				Id:           id,
				Email:        fake.Email(),
				Nickname:     nick,
				HashPassword: hash,
				Score:        rand.Int(),
				AvatarLink:   DefaultAvatarLink,
			}

		}

	})
}

func getUsers() map[string]User {
	mu.Lock()
	fmt.Println(users)
	mu.Unlock()

	return users
}

func GetUser(login string) (User, error) {
	defer mu.Unlock()

	mu.Lock()
	if _, ok := users[login]; !ok {
		return User{}, errors.New("Invalid login")
	} else {
		return users[login], nil
	}
}

func findUserById(id int64) (User, error) {
	defer mu.Unlock()

	mu.Lock()
	for _, val := range users {
		if val.Id == id {
			return val, nil
		}
	}

	return User{}, errors.New("user with this id not found")
}

func addUser(in User) error {
	defer mu.Unlock()
	mu.Lock()

	if _, ok := users[in.Nickname]; ok {
		return errors.New("User with this nickname already exist")
	}

	users[in.Nickname] = in

	return nil
}

func PrintUsers() {
	mu.Lock()

	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")

	mu.Unlock()
}

func getNextId() int64 {
	atomic.AddInt64(&currentId, 1)

	return currentId
}

func updateDBUser(user User) error {
	defer mu.Unlock()

	mu.Lock()
	if _, ok := users[user.Nickname]; !ok {
		return errors.New("cannot find user")
	}

	users[user.Nickname] = user
	return nil
}

type ByNameScore []User

func (a ByNameScore) Len() int           { return len(a) }
func (a ByNameScore) Less(i, j int) bool { return a[i].Score < a[j].Score }
func (a ByNameScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getScores() []User {
	mu.Lock()
	result := make([]User, 0, 1)
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
