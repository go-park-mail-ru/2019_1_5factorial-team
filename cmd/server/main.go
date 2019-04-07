package main

import (
	"flag"
	"fmt"
	"log"

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

	s := server.New(*port)
	err = s.Run()
	if err != nil {
		panic(err)
	}
}
