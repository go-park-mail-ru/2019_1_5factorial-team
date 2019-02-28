package server

import (
	"2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Run(portIn string) error {
	if portIn == "" {
		portIn = "5051"
	}
	portFull := ":" + portIn

	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HW).Methods("GET")

	return http.ListenAndServe(portFull, router)
}
