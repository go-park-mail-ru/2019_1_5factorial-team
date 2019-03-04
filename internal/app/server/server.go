package server

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Run(port string) error {
	address := ":" + port

	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")

	return http.ListenAndServe(address, router)
}
