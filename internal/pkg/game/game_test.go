package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/sirupsen/logrus"
	"testing"
)

func InitConfig() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	log.InitLogs()
}

func TestNewPlayer(t *testing.T) {
	InitConfig()
	p := NewPlayer(nil, "kek")
	if p.Score != 0 || p.Token != "kek" {
		t.Error("wrong create player")
	}

	go panicWorker.PanicWorker(p.Listen)
	p.CloseConn()
}

func TestNewRoom(t *testing.T) {
	InitConfig()
	var maxp uint = 10
	game := NewGame(10, nil)

	r := NewRoom(maxp, game)
	if len(r.ID) == 0 {
		t.Error("empty room ID")
	}

	if r.MaxPlayers != maxp {
		t.Error("wrong max players count")
	}

	go r.Run()

	r.Close()
}
