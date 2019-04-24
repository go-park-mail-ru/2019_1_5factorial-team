package chat

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/chat"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

//var InstanceChat *MyGorgeousChat

type MyGorgeousChat struct {
	//InstanceChat *chat.Chat
	port         string
}

func New(port string) *MyGorgeousChat {
	mgc := MyGorgeousChat{}
	mgc.port = port

	//mgc.InstanceChat = chat.NewChat(20)
	//go panicWorker.PanicWorker(mgc.InstanceChat.Start)

	//InstanceChat = &mgc
	chat.Start()

	return &mgc
}

func (mgc *MyGorgeousChat) Run() error {

	address := ":" + mgc.port

	chatRouter := mux.NewRouter()
	chatRouter.Use(middleware.CORSMiddleware)
	chatRouter.Use(middleware.AuthMiddleware)
	chatRouter.Use(middleware.CheckLoginMiddleware)
	chatRouter.Use(middleware.PanicMiddleware)

	chatRouter.HandleFunc("/api/chat/global/ws", controllers.ConnectToGlobalChat).Methods("GET", "OPTIONS")

	err := http.ListenAndServe(address, chatRouter)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
