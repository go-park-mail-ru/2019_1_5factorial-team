package donation

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

type DonateNotification struct {
	port string
}

func New(port string) *DonateNotification {
	mgc := DonateNotification{}
	mgc.port = port

	return &mgc
}

func (mgg *DonateNotification) Run() error {
	address := ":" + mgg.port
	log.Println(address)

	router := mux.NewRouter()
	//router.Use(middleware.CORSMiddleware)
	router.Use(middleware.PanicMiddleware)

	router.HandleFunc("/api/push", controllers.GetNotification)
	//gameRouter.HandleFunc("/api/game/ws", controllers.Play).Methods("GET", "OPTIONS")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}