package controllers

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/chat"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ConnectToGlobalChat(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)

	var user *chat.User
	// TODO(smet1): ping-pong or ws will be close on timeout
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

	if req.Context().Value("authorized").(bool) != false {
		user, err = chat.NewUserID(conn, req.Context().Value("userID").(string), authGRPC)
		if err != nil {
			ErrResponse(res, http.StatusInternalServerError, "cant connect to chat")

			ctxLogger.Error(errors.Wrap(err, "cant create chat.User"))
			return
		}
	} else {
		user, err = chat.NewUserFake(conn)
		if err != nil {
			ErrResponse(res, http.StatusInternalServerError, "cant connect to chat")

			ctxLogger.Error(errors.Wrap(err, "cant create chat.User"))
			return
		}
	}

	if user != nil {
		go panicWorker.PanicWorker(user.Listen)
		chat.InstanceChat.AddUser(user)
	}
}
