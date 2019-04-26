package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/chat"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
)

func main() {
	// TODO(): кушать только необходимые коннекшены к базам

	port := flag.String("port", "5052", "chat-server will start on this port")
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	flag.Parse()

	log.Warn("chat-server will start on port ", *port)
	log.Warn("config path: ", *configPath)

	err := config.Init(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.InitLogs()

	database.InitConnection()

	c := chat.New(*port)
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
