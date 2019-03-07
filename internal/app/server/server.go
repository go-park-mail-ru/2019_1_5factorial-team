package server

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func Run(port string) error {
	address := ":" + port
	router := mux.NewRouter()

	// TODO: CORS, panic and auth middleware
	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/hello", controllers.HW).Methods("GET")
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/session", controllers.SignIn).Methods("POST")
	router.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE")
	router.HandleFunc("/api/session", controllers.GetUserFromSession).Methods("GET")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
