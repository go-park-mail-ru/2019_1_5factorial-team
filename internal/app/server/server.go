package server

import (
	"fmt"
	_ "github.com/go-park-mail-ru/2019_1_5factorial-team/docs"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type MyGorgeousServer struct {
	port string
}

func New(port string) *MyGorgeousServer {
	mgs := MyGorgeousServer{}
	mgs.port = port

	database.InitConnection()

	// TODO(): говно - переделать
	// хочу передавать в json'е имя функи для фейк генерации и здесь ее вызывать
	// avoid cycle imports
	for _, conn := range config.Get().DBUserConfig {
		if conn.GenerateFakeUsers {
			fu := user.GenerateUsers()

			for i, val := range fu {
				fmt.Println(i, "| id:", val.ID.Hex(), ", Nick:", val.Nickname, ", Password:", val.Nickname)

				col, err := database.GetCollection(conn.CollectionName)
				if err != nil {
					log.Fatal(errors.Wrap(err, "collection empty"))
				}

				err = col.Insert(val)
				if err != nil {
					log.Fatal(errors.Wrap(err, "error while adding new user"))
				}
			}
		}
	}

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
