package user

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/manveru/faker"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

func GenerateUsers() []User {
	fmt.Println("---=== GENERATE FAKE USERS IN PROGRESS ===---")
	u := make([]User, 0, config.Get().FakeUsersConfig.UsersCount)
	fake, _ := faker.New(config.Get().FakeUsersConfig.Lang)
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

	for i := 0; i < config.Get().FakeUsersConfig.UsersCount; i++ {
		nick := fake.FirstName()
		hash, _ := GetPasswordHash(nick)

		u = append(u, User{
			ID:           bson.NewObjectId(),
			Email:        fake.Email(),
			Nickname:     nick,
			HashPassword: hash,
			Score:        rand.Intn(config.Get().FakeUsersConfig.MaxScore),
			AvatarLink:   "",
		})
	}

	fmt.Println("---=== GENERATE FAKE USERS DONE ===---")

	return u
}
