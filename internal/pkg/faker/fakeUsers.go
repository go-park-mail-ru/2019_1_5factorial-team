package faker

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"log"
)

type FakeUsersConfig struct {
	UsersCount int `json:"users_count"`
	Lang string `json:"lang"`
	MaxScore int `json:"max_score"`
}

var FakeUsersConf = FakeUsersConfig{}

func init() {
	err := config_reader.ReadConfigFile("user_faker_config.json", &FakeUsersConf)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading User faker config"))
	}
	fmt.Println(FakeUsersConf)

}

func GenerateUsers() []user.User {
	fmt.Println("---=== GENERATE FAKE USERS IN PROGRESS ===---")
	//TODO(): доделать генерацию
	fmt.Println("---=== GENERATE FAKE USERS DONE ===---")

}
