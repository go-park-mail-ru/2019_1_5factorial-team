package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/game"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Play(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		ctxLogger.Printf("error while connecting: %s", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		ctxLogger.Error(errors.Wrap(err, "not authorized"))
		return
	}

	ctxLogger.Warn(err)

	player := game.NewPlayer(conn, currentSession.Value)
	go panicWorker.PanicWorker(player.Listen)
	game.InstanceGame.AddPlayer(player)
}

func CreateUniqueRoom(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		ctxLogger.Printf("error while connecting: %s", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		ctxLogger.Error(errors.Wrap(err, "not authorized"))
		return
	}

	ctxLogger.Warn(err)

	player := game.NewPlayer(conn, currentSession.Value)
	go panicWorker.PanicWorker(player.Listen)
	game.InstanceGame.AddUniquePlayer(player)
}

func ConnectRoomByLink(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	query := req.URL.Query()

	ctxLogger.Info("query = ", query)

	roomID := query.Get("room")
	if !game.InstanceGame.CheckRoomLink(roomID) {
		ErrResponse(res, http.StatusBadRequest, "room not found or full")

		ctxLogger.Error(errors.New("invalid room ID"))
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)

		ctxLogger.Errorf("error while connecting: %s", err.Error())
		return
	}

	currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		ctxLogger.Error(errors.Wrap(err, "not authorized"))
		return
	}

	ctxLogger.Warn(err)

	player := game.NewPlayer(conn, currentSession.Value)
	go panicWorker.PanicWorker(player.Listen)
	game.InstanceGame.AddUniquePlayerLink(player, roomID)
}
