package server

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func Run(port string) error {
	address := ":" + port
	router := mux.NewRouter()

	// TODO: CORS, panic and auth middleware

	router.HandleFunc("/hello", controllers.HW).Methods("GET")
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}