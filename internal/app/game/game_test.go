package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"testing"
)

var casesNew = []struct {
	port string
	want string
}{
	{"1", "1"},
	{"", ""},
}

func TestNew(t *testing.T) {
	for _, val := range casesNew {
		s := New(val.port)
		if s.port != val.port {
			t.Error("ERROR expected:", val.port, "have:", s.port)
		}
	}
}

func TestMyGorgeousGame_Run(t *testing.T) {
	err := config.Init("/etc/5factorial/")
	if err != nil {
		log.Fatal(err.Error())
	}

	config.Get().DBConfig[0].Hostname = "localhost"
	config.Get().DBConfig[0].MongoPort = "27061"
	config.Get().DBConfig[0].TruncateTable = true

	config.Get().DBConfig[1].Hostname = "localhost"
	config.Get().DBConfig[1].MongoPort = "27062"
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

	s := New("aaaa")
	err = s.Run()
	if err == nil {
		t.Error("error expected")
	}
}
