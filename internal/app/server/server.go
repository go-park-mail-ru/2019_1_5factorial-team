package server

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func Run(port string) error {
	address := ":" + port
	router := mux.NewRouter()

	// TODO: panic
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.CORSMiddleware)

	router.HandleFunc("/hello", controllers.HW).Methods("GET")
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/{id:[0-9]+}", controllers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/session", controllers.SignIn).Methods("POST", "OPTIONS")

	routerLoginRequired := router.PathPrefix("").Subrouter()

	routerLoginRequired.Use(middleware.CheckLoginMiddleware)

	routerLoginRequired.HandleFunc("/api/user", controllers.GetUserFromSession).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/user", controllers.UpdateProfile).Methods("PUT", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.IsSessionValid).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/docs", httpSwagger.WrapHandler)

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
