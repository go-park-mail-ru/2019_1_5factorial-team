package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/fileproc"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Run(port string) error {

	address := ":" + port
	router := mux.NewRouter()

	// TODO: CORS, panic and auth middleware
	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/upload_avatar", controllers.UploadAvatar).Methods("POST")
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir(fileproc.UploadPath)))
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/session", controllers.SignIn).Methods("POST")

	routerLoginRequired := router.PathPrefix("").Subrouter()

	routerLoginRequired.Use(middleware.CheckLoginMiddleware)

	routerLoginRequired.HandleFunc("/api/user", controllers.GetUserFromSession).Methods("GET")
	routerLoginRequired.HandleFunc("/api/session", controllers.IsSessionValid).Methods("GET")
	routerLoginRequired.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
