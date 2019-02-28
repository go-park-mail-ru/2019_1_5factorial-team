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
	router.HandleFunc("/user/register", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}