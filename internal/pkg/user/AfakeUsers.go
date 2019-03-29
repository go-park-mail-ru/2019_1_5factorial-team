package user

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/manveru/faker"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

//type FakeUsersConfig struct {
//	UsersCount int    `json:"users_count"`
//	Lang       string `json:"lang"`
//	MaxScore   int    `json:"max_score"`
//}
//
//var FakeUsersConf = FakeUsersConfig{}
//
//func init() {
//	err := config_reader.ReadConfigFile("user_faker_config.json", &FakeUsersConf)
//	if err != nil {
//		log.Fatal(errors.Wrap(err, "error while reading User faker config"))
//	}
//	fmt.Println("Faker init:", FakeUsersConf)
//
//}

func GenerateUsers() []DatabaseUser {
	fmt.Println("---=== GENERATE FAKE USERS IN PROGRESS ===---")
	u := make([]DatabaseUser, 0, config.GetInstance().FakeUsersConfig.UsersCount)
	fake, _ := faker.New(config.GetInstance().FakeUsersConfig.Lang)
	fake.Rand = rand.New(rand.NewSource(42))

	// наш самый любимый юзер, с истоков нашего проекта
	hash, _ := GetPasswordHash("password")
	u = append(u, DatabaseUser{
		CollectionID: bson.NewObjectId(),
		Email:        "kek.k.ek",
		Nickname:     "kekkekkek",
		HashPassword: hash,
		Score:        rand.Intn(config.GetInstance().FakeUsersConfig.MaxScore),
		AvatarLink:   "",
	})

	for i := 0; i < config.GetInstance().FakeUsersConfig.UsersCount; i++ {
		nick := fake.FirstName()
		hash, _ := GetPasswordHash(nick)

		u = append(u, DatabaseUser{
			CollectionID: bson.NewObjectId(),
			Email:        fake.Email(),
			Nickname:     nick,
			HashPassword: hash,
			Score:        rand.Intn(config.GetInstance().FakeUsersConfig.MaxScore),
			AvatarLink:   "",
		})
	}

	fmt.Println("---=== GENERATE FAKE USERS DONE ===---")

	return u
}
