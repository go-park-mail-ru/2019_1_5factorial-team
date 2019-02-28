package main

import (
	"2019_1_5factorial-team/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	port := ":5051"
	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HW).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}