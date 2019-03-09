package server

import (
	"flag"
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var fileServerDir = flag.String("fileServerDir", "/var/www/media/factorial", "way to static files")

// var fileServerDir = flag.String("fileServerDir", "../../src", "way to static files")
func Run(port string) error {

	address := ":" + port
	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/upload_avatar", controllers.UploadAvatar).Methods("POST")
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir(*fileServerDir)))

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
