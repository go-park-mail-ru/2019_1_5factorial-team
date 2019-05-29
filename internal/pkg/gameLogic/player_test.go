package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func InitBasics() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	config.Get().DBConfig[0].Hostname = "localhost"
	config.Get().DBConfig[0].MongoPort = "27091"
	config.Get().DBConfig[0].TruncateTable = true

	config.Get().DBConfig[1].Hostname = "localhost"
	config.Get().DBConfig[1].MongoPort = "27092"
	config.Get().DBConfig[1].TruncateTable = true

	config.Get().DBConfig[2].Hostname = "localhost"
	config.Get().DBConfig[2].MongoPort = "27063"
	config.Get().DBConfig[2].TruncateTable = true

	config.Get().AuthGRPCConfig.Hostname = "localhost"

	log.InitLogs()

	database.InitConnection()

	err = gRPC.InitAuthClient()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestNewPlayerCharacter(t *testing.T) {
	InitBasics()
	authGRPC := grpcAuth.AuthGRPCClient
	if authGRPC == nil {
		t.Error("authGRPC shoulnt be nil")
		return
	}

	_, err := NewPlayerCharacter("kek", authGRPC)
	if err == nil {
		t.Error("should be err bcs of invalid token")
		return
	}
}
