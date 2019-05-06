package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/game"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

const roomsCount uint32 = 10

type MyGorgeousGame struct {
	port string
}

func New(port string) *MyGorgeousGame {
	mgc := MyGorgeousGame{}
	mgc.port = port

	game.Start(roomsCount, grpcAuth.AuthGRPCClient)

	return &mgc
}

func (mgg *MyGorgeousGame) Run() error {

	address := ":" + mgg.port

	gameRouter := mux.NewRouter()
	gameRouter.Use(middleware.CORSMiddleware)
	gameRouter.Use(middleware.AuthMiddleware)
	gameRouter.Use(middleware.CheckLoginMiddleware)
	gameRouter.Use(middleware.PanicMiddleware)

	gameRouter.HandleFunc("/api/game/ws", controllers.Play).Methods("GET", "OPTIONS")

	err := http.ListenAndServe(address, gameRouter)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
