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

	// TODO: CORS, panic
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.CORSMiddleware)

	router.HandleFunc("/hello", controllers.HW).Methods("GET")
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/session", controllers.SignIn).Methods("POST", "OPTIONS")

	routerLoginRequired := router.PathPrefix("").Subrouter()

	routerLoginRequired.Use(middleware.CheckLoginMiddleware)

	routerLoginRequired.HandleFunc("/api/user", controllers.GetUserFromSession).Methods("GET")
	routerLoginRequired.HandleFunc("/api/user", controllers.UpdateProfile).Methods("PUT")
	routerLoginRequired.HandleFunc("/api/session", controllers.IsSessionValid).Methods("GET")
	routerLoginRequired.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
