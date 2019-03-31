package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/server"
)

func main() {
	port := flag.String("port", "5051", "server will start on this port")
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	flag.Parse()

	fmt.Println("server will start on port", *port)
	fmt.Println("config path:", *configPath)

	err := config.Init(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// мб вынести саму базу в app?
	database.InitConnection()

	// вот в душе не знаю куда это засунуть, ибо если это оставить в internal/pkg/database/database.go
	// что впринципе логично, возникает мои любимые циклические конфликты
	if config.Get().DBUserConfig.GenerateFakeUsers {
		fu := user.GenerateUsers()

		for i, val := range fu {
			fmt.Println(i, "| id:", val.ID.Hex(), ", Nick:", val.Nickname, ", Password:", val.Nickname)

			err := database.GetUserCollection().Insert(val)
			if err != nil {
				log.Fatal(errors.Wrap(err, "error while adding new user"))
			}
		}
	}

	s := server.New(*port)
	err = s.Run()
	if err != nil {
		panic(err)
	}
}
