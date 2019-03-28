package server

import (
	"fmt"
	_ "github.com/go-park-mail-ru/2019_1_5factorial-team/docs"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/fileproc"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"

	"github.com/swaggo/http-swagger"
	"net/http"
)

var instance *MyGorgeousServer

type MyGorgeousServer struct {
	StaticServerConfig fileproc.StaticServerConfig
	CORSConfig         middleware.CORSConfig
	CookieConfig       session.CookieConfig

	configPath string
}

func (mgs *MyGorgeousServer) New(config string) *MyGorgeousServer {

	mgs.configPath = config

	err := config_reader.ReadConfigFile(config, "static_server_config.json", &mgs.StaticServerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading static_server_config config"))
	}
	mgs.StaticServerConfig.MaxUploadSize = mgs.StaticServerConfig.MaxUploadSizeMB * 1024 * 1024

	log.Println("New Server->Static server config = ", mgs.StaticServerConfig)


	err = config_reader.ReadConfigFile(config, "cors_config.json", &mgs.CORSConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading CORS config"))
	}

	log.Println("New Server->CORS config = ", mgs.CORSConfig)


	err = config_reader.ReadConfigFile(config, "cookie_config.json", &mgs.CookieConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	}
	fmt.Println("New Server->Cookie config = ", mgs.CookieConfig)

	// чтобы не копипастить везде (это время жизни куки в часах, чтобы добавлять к уже созданной)
	//CookieConf.CookieTimeHours = time.Duration(CookieConf.CookieDuration * int64(time.Hour))

	instance = mgs

	return mgs
}

func Run(port string) error {

	address := ":" + port
	router := mux.NewRouter()

	// TODO: panic
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.CORSMiddleware)

	router.HandleFunc("/hello", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/user", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id:[0-9]+}", controllers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/score", controllers.GetLeaderboard).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/count", controllers.UsersCountInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/session", controllers.SignIn).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/upload_avatar", controllers.UploadAvatar).Methods("POST", "OPTIONS")

	imgServer := http.FileServer(http.Dir(fileproc.StaticConfig.UploadPath))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", imgServer))

	routerLoginRequired := router.PathPrefix("").Subrouter()

	routerLoginRequired.Use(middleware.CheckLoginMiddleware)

	routerLoginRequired.HandleFunc("/api/user", controllers.GetUserFromSession).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/user", controllers.UpdateProfile).Methods("PUT", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.IsSessionValid).Methods("GET", "OPTIONS")
	routerLoginRequired.HandleFunc("/api/session", controllers.SignOut).Methods("DELETE", "OPTIONS")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	err := http.ListenAndServe(address, router)
	if err != nil {
		return errors.Wrap(err, "server Run error")
	}

	return nil
}
