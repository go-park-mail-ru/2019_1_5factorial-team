package server

import (
	_ "github.com/go-park-mail-ru/2019_1_5factorial-team/docs"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rossmcdonald/telegram_hook"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	"net/http"
	"os"
)

type MyGorgeousServer struct {
	port string
}

func InitLogs()  {
	// настраиваем logrus (по всему проекту)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   config.Get().LogrusConfig.DisableColors,
		FullTimestamp:   config.Get().LogrusConfig.FullTimestamp,
		TimestampFormat: config.Get().LogrusConfig.TimestampFormat,
	})
	//log.SetFormatter(&log.JSONFormatter{
	//	TimestampFormat: config.Get().LogrusConfig.TimestampFormat,
	//	PrettyPrint: true,
	//})

	if config.Get().LogrusConfig.AppName != "" {
		// тележка <3
		hook, err := telegram_hook.NewTelegramHook(
			config.Get().LogrusConfig.AppName,
			config.Get().LogrusConfig.AuthToken,
			config.Get().LogrusConfig.TargetID,
			telegram_hook.WithAsync(config.Get().LogrusConfig.Async),
			telegram_hook.WithTimeout(config.Get().LogrusConfig.Timeout.Duration),
		)
		if err != nil {
			log.Fatalf("Encountered error when creating Telegram hook: %s", err)
		}
		log.AddHook(hook)
	}
}

func New(port string) *MyGorgeousServer {
	InitLogs()

	mgs := MyGorgeousServer{}
	mgs.port = port

	database.InitConnection()

	return &mgs
}

func (mgs *MyGorgeousServer) Run() error {

	address := ":" + mgs.port
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

	imgServer := http.FileServer(http.Dir(config.Get().StaticServerConfig.UploadPath))
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
