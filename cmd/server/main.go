package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/server"
	"os"
)

var defaultPort = "5051"

func main() {
	var port string
	if len(os.Args) == 1 {
		port = defaultPort
	} else {
		port = os.Args[1]
	}

	fmt.Println("server will start on port", port)

	err := server.Run(port)
	if err != nil {
		fmt.Println("error happened: ", err)
	}
}
