package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/donation"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
)

func main() {
	port := flag.String("port", "5054", "donation-server will start on this port")
	configPath := flag.String("config", "/etc/5factorial/donation/", "dir with server configs")
	flag.Parse()

	log.InitLogs()

	log.Warn("donation-server will start on port ", *port)
	log.Warn("config path: ", *configPath)

	err := config.Init(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	c := donation.New(*port)
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
