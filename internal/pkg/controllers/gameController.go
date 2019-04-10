package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/game"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func Play(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	ctxLogger.Info("============================================")

	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(res, req, http.Header{"Upgrade": []string{"websocket"}})
	if err != nil {
		ctxLogger.Printf("error while connecting: %s", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print("connected to client")

	player := game.NewPlayer(conn, req.Context().Value("userID").(string))
	go player.Listen()
	game.InstanceGame.AddPlayer(player)

	log.Print("Play exit")
}
