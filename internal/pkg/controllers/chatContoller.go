package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/chat"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ConnectToGlobalChat(res http.ResponseWriter, req *http.Request) {
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

	//currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	//if err == http.ErrNoCookie {
	//	// бесполезная проверка, так кука валидна, но по гостайлу нужна
	//	ErrResponse(res, http.StatusUnauthorized, "not authorized")
	//
	//	ctxLogger.Error(errors.Wrap(err, "not authorized"))
	//	return
	//}
	//
	//user, err := chat.NewUser(conn, currentSession.Value)
	//if err != nil {
	//	ErrResponse(res, http.StatusInternalServerError, "cant connect to chat")
	//
	//	ctxLogger.Error(errors.Wrap(err, "cant create chat.User"))
	//	return
	//}

	if req.Context().Value("authorized").(bool) != false {
		user, err := chat.NewUserID(conn, req.Context().Value("userID").(string))
		if err != nil {
			ErrResponse(res, http.StatusInternalServerError, "cant connect to chat")

			ctxLogger.Error(errors.Wrap(err, "cant create chat.User"))
			return
		}
		// TODO(): убрать
		go panicWorker.PanicWorker(user.Listen)
		chat.InstanceChat.AddUser(user)
	} else {
		user, err := chat.NewUserFake(conn)
		if err != nil {
			ErrResponse(res, http.StatusInternalServerError, "cant connect to chat")

			ctxLogger.Error(errors.Wrap(err, "cant create chat.User"))
			return
		}

		go panicWorker.PanicWorker(user.Listen)
		chat.InstanceChat.AddUser(user)
	}
	//go panicWorker.PanicWorker(user.Listen)
	//chat.InstanceChat.AddUser(user)
	//chat2.InstanceChat.InstanceChat.AddUser(user)
}
