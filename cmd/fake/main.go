package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// заполняем базу user, коллекция profile данными
func main() {
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	port := flag.String("port", "27031", "user db port")

	flag.Parse()

	fuc := user.FakeUsersConfig{}
 
	// конфиг генерации фейковых юзеров
	err := config_reader.ReadConfigFile(*configPath, "user_faker_config.json", &fuc)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading user faker config"))
	}
	log.Println("User faker config = ", fuc)

	fmt.Println(fuc)
	fmt.Println(*port)

	fu := user.GenerateUsers(fuc)

	session, err := mgo.Dial("mongodb://localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}

	collection := session.DB("user").C("profile")

	for i, val := range fu {
		fmt.Println(i, "| id:", val.ID.Hex(), ", Nick:", val.Nickname, ", Password:", val.Nickname)

		err = collection.Insert(val)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error while adding new user"))
		}
	}

	session.Close()
}
