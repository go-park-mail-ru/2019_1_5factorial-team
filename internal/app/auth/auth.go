package auth

import (
	"flag"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
)

func Run()  {
	port := flag.String("port", "5000", "chat-server will start on this port")
	configPath := flag.String("config", "/etc/5factorial/", "dir with server configs")
	flag.Parse()

	log.Warn("auth-service will start on port ", *port)
	log.Warn("config path: ", *configPath)

	err := config.Init(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.InitLogs()

	database.InitConnection()

	err = session.GRPCServer()
	if err != nil {
		log.Error(errors.Wrap(err, "cant start auth service"))
	}
}
