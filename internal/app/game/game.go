package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/game"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type MyGorgeousGame struct {
	port string
}

func New(port string) *MyGorgeousGame {
	mgc := MyGorgeousGame{}
	mgc.port = port

	game.Start(config.Get().GameConfig.MaxRooms, grpcAuth.AuthGRPCClient)

	return &mgc
}

func (mgg *MyGorgeousGame) Run() error {
	address := ":" + mgg.port
	log.Println(address)

	router := mux.NewRouter()
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.PanicMiddleware)
	router.Path("/metrics").Handler(promhttp.Handler())

	gameRouter := router.PathPrefix("").Subrouter()
	//gameRouter.Use(middleware.CheckStatus)
	gameRouter.Use(middleware.AuthMiddleware)
	gameRouter.Use(middleware.CheckLoginMiddleware)

	gameRouter.HandleFunc("/api/game/ws", controllers.Play).Methods("GET", "OPTIONS")
	gameRouter.HandleFunc("/api/game/friend", controllers.CreateUniqueRoom).Methods("GET", "OPTIONS")
	gameRouter.HandleFunc("/api/game/connect", controllers.ConnectRoomByLink).Methods("GET", "OPTIONS")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
