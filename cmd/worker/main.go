package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/pkg/errors"
)

// заполняем базу user, коллекция profile данными
func main() {
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	port := flag.String("port", "27032", "user db port")

	flag.Parse()

	minuteSleep := session.WorkerSession{}

	// конфиг для воркера
	err := config_reader.ReadConfigFile(*configPath, "session_worker_config.json", &minuteSleep)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading session worker config"))
	}
	log.Println("Session worker config = ", minuteSleep)

	fmt.Println(minuteSleep)
	fmt.Println(*port)

	go session.RemoveSession(port, minuteSleep.SleepTime)

}
