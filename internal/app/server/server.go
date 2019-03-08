package server

import (
	"flag"
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var fileServerDir = flag.String("fileServerDir", "../../src", "way to static files")

// TODO:  вынести src
// TODO: добавить константный путь
//var root = flag.String("root", "./src/img", "file system path")

func Run(port string) error {
	address := ":" + port
	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/getAvatar", controllers.GetImg).Methods("GET")
	router.HandleFunc("/uploadAvatar", controllers.UploadAvatar).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(*fileServerDir))) // работает не трогать </img/avatar.png>

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
