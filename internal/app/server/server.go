package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Run(port string) error {
	address := ":" + port
	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")

	router.HandleFunc("/getimg", controllers.GetImg).Methods("GET")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
