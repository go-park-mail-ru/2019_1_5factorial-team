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
	"time"
)

type MyGorgeousServer struct {
	port string
}

func New(port string) *MyGorgeousServer {
	// настраиваем logrus (по всему проекту)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//httpTransport := &http.Transport{}
	//httpClient := &http.Client{Transport: httpTransport}
	//dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	//httpTransport.Dial = dialer.Dial
	//
	//hook, err := telegram_hook.NewTelegramHookWithClient(
	//	"5factorial",
	//	"871491595:AAEpe6PSwbbV96dpeUSiugpkhQs-jCd0hCg",
	//	"149677494",
	//	httpClient,
	//	telegram_hook.WithAsync(true),
	//	telegram_hook.WithTimeout(30 * time.Second),
	//)

	hook, err := telegram_hook.NewTelegramHook(
		"5factorial",
		"871491595:AAEpe6PSwbbV96dpeUSiugpkhQs-jCd0hCg",
		"149677494",
		telegram_hook.WithAsync(true),
		telegram_hook.WithTimeout(30 * time.Second),
	)
	if err != nil {
		log.Fatalf("Encountered error when creating Telegram hook: %s", err)
	}
	log.AddHook(hook)

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
