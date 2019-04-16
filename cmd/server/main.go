package main

import (
	"flag"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/server"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
)

func main() {
	port := flag.String("port", "5051", "server will start on this port")
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	flag.Parse()

	log.Warn("server will start on port", *port)
	log.Warn("config path:", *configPath)

	err := config.Init(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.InitLogs()

	database.InitConnection()

	s := server.New(*port)

	go session.CleanSession()

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
