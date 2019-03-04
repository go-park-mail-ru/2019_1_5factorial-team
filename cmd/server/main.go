package main

import (
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/server"
)


func main() {
	port := flag.String("port", "5051", "server will start on this port")
	flag.Parse()

	fmt.Println("server will start on port", *port)

	err := server.Run(*port)
	if err != nil {
		fmt.Println("error happened: ", err)
	}
}
