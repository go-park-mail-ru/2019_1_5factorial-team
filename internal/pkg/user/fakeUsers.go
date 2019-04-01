package user

import (
	"fmt"
	"github.com/manveru/faker"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

// структура конфига генератора фейковых юзеров
type FakeUsersConfig struct {
	UsersCount int    `json:"users_count"`
	Lang       string `json:"lang"`
	MaxScore   int    `json:"max_score"`
}

func GenerateUsers(config FakeUsersConfig) []User {
	fmt.Println("---=== GENERATE FAKE USERS IN PROGRESS ===---")
	u := make([]User, 0, config.UsersCount)
	fake, _ := faker.New(config.Lang)
	fake.Rand = rand.New(rand.NewSource(42))

	// наш самый любимый юзер, с истоков нашего проекта
	hash, _ := GetPasswordHash("password")
	u = append(u, User{
		ID:           bson.NewObjectId(),
		Email:        "kek.k.ek",
		Nickname:     "kekkekkek",
		HashPassword: hash,
		Score:        100500,
		AvatarLink:   "",
	})

	for i := 0; i < config.UsersCount; i++ {
		nick := fake.FirstName()
		hash, _ := GetPasswordHash(nick)

		u = append(u, User{
			ID:           bson.NewObjectId(),
			Email:        fake.Email(),
			Nickname:     nick,
			HashPassword: hash,
			Score:        rand.Intn(config.MaxScore),
			AvatarLink:   "",
		})
	}

	fmt.Println("---=== GENERATE FAKE USERS DONE ===---")

	return u
}
