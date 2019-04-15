package server

import (
	"net/http"

	_ "github.com/go-park-mail-ru/2019_1_5factorial-team/docs"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type MyGorgeousServer struct {
	port string
}

func New(port string) *MyGorgeousServer {
	mgs := MyGorgeousServer{}
	mgs.port = port

	return &mgs
}

func (mgs *MyGorgeousServer) Run() error {

	address := ":" + mgs.port
	router := mux.NewRouter()
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.PanicMiddleware)

	// routers for sharing static and swagger
	staticRouter := router.PathPrefix("").Subrouter()
	imgServer := http.FileServer(http.Dir(config.Get().StaticServerConfig.UploadPath))
	staticRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", imgServer))
	staticRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// main router with check cookie
	mainRouter := router.PathPrefix("").Subrouter()
	mainRouter.Use(middleware.AuthMiddleware)

	mainRouter.HandleFunc("/api/session/oauth/google", controllers.LoginFromGoogle).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/session/oauth/vk", controllers.LoginFromVK).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/session/oauth/yandex", controllers.LoginFromYandex).Methods("POST", "OPTIONS")

	mainRouter.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")
	mainRouter.HandleFunc("/api/user", controllers.SignUp).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/user/{id:[0-9]+}", controllers.GetUser).Methods("GET", "OPTIONS")
	mainRouter.HandleFunc("/api/user/score", controllers.GetLeaderboard).Methods("GET", "OPTIONS")
	mainRouter.HandleFunc("/api/user/count", controllers.UsersCountInfo).Methods("GET", "OPTIONS")
	mainRouter.HandleFunc("/api/session", controllers.SignIn).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/upload_avatar", controllers.UploadAvatar).Methods("POST", "OPTIONS")

	// routers with logged user
	routerLoginRequired := mainRouter.PathPrefix("").Subrouter()
	routerLoginRequired.Use(middleware.CheckLoginMiddleware)

	routerLoginRequired.HandleFunc("/api/user", controllers.GetUserFromSession).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/user", controllers.UpdateProfile).Methods("PUT", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.IsSessionValid).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE", "OPTIONS")

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
