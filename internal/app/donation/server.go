package donation

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/donation"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type MyFabulousDonation struct {
	port string
}

func New(port string) *MyFabulousDonation {
	donation := MyFabulousDonation{}
	donation.port = port

	return &donation
}

func (mgg *MyFabulousDonation) Run() error {
	address := ":" + mgg.port
	log.Println(address)

	donationInstance := donation.NewServer()
	err := donationInstance.Run()
	if err != nil {
		return errors.Wrap(err, "failed to init yaClient")
	}

	router := mux.NewRouter()
	router.Use(middleware.PanicMiddleware)

	router.HandleFunc("/api/donation/ws", donationInstance.GetWssHandler()).Methods("GET", "OPTIONS") // TODO: implementirovat'

	router.HandleFunc("/api/donation/notification", donationInstance.GetNotificationHandler()).Methods("POST", "OPTIONS")

	err = http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
