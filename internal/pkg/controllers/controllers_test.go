package controllers

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCases struct {
	routerPath     string
	method         string
	url            string
	body           io.Reader
	urlValues      string
	expectedRes    string
	expectedStatus int
	authCtx        bool
	userIDCtx      string
}

var funcs = []func(*testing.T){
	TestSignUp,
	//TestGetUserFromSession,
	TestGetLeaderboard,
	TestUsersCountInfo,
	//TestGetUser,
	TestSignIn,
}

func TestControllers(t *testing.T) {
	MainInit()
	_, _ = user.CreateUser("kekkekkek", "kek.k.ek", "password", "")
	//for _, test := range funcs {
	//	test(t)
	//}
}

func testHandler(funcToTest func(http.ResponseWriter, *http.Request), tests []TestCases, t *testing.T) error {
	var req *http.Request
	var err error
	for i, val := range tests {
		logrus.Warnf("test #%d, val = %#v", i, val)

		req, err = http.NewRequest(val.method, val.url, val.body)
		if err != nil {
			return err
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc(val.routerPath, funcToTest).Methods(val.method)
		authGRPC := grpcAuth.AuthGRPCClient

		ctx := req.Context()
		ctx = context.WithValue(ctx, "authorized", val.authCtx)
		ctx = context.WithValue(ctx, "userID", val.userIDCtx)
		ctx = context.WithValue(ctx, "authGRPC", authGRPC)
		if val.authCtx {
			ctx = context.WithValue(ctx, "logger", log.LoggerWithAuth(req.WithContext(ctx)))
		} else {
			ctx = context.WithValue(ctx, "logger", log.LoggerWithoutAuth(req.WithContext(ctx)))
		}
		req = req.WithContext(ctx)

		router.ServeHTTP(rr, req)

		// try to check cookie, LOOKS LIKE HACK
		//kek := rr.Header()["Set-Cookie"]
		//if kek != nil {
		//	log.Println(strings.Split(kek[0], ";"))
		//}

		// Check the status code is what we expect.
		if status := rr.Code; status != val.expectedStatus {
			t.Errorf("handler returned wrong status code: \ngot  %v \nwant %v\ntest #%d:\n\t%v\n",
				status, val.expectedStatus, i, val)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			t.Errorf("handler returned unexpected body: \ngot  %v \nwant %v\ntest #%d:\n\t%v\n",
				rr.Body.String(), val.expectedRes, i, val)
		}
	}
	return nil
}

func MainInit() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	config.Get().DBConfig[0].Hostname = "localhost"
	config.Get().DBConfig[0].MongoPort = "27061"
	config.Get().DBConfig[0].TruncateTable = true

	config.Get().DBConfig[1].Hostname = "localhost"
	config.Get().DBConfig[1].MongoPort = "27062"
	config.Get().DBConfig[1].TruncateTable = true

	config.Get().DBConfig[2].Hostname = "localhost"
	config.Get().DBConfig[2].MongoPort = "27063"
	config.Get().DBConfig[2].TruncateTable = true

	config.Get().AuthGRPCConfig.Hostname = "localhost"

	log.InitLogs()

	database.InitConnection()

	err = gRPC.InitAuthClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// indexes user
	col, _ := database.GetCollection(config.Get().DBConfig[0].CollectionName)
	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"nickname"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// indexes session
	col, _ = database.GetCollection(config.Get().DBConfig[1].CollectionName)
	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"user_id"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"token"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// генерим юзеров на базенке
	fuc := user.FakeUsersConfig{}

	// конфиг генерации фейковых юзеров
	_ = config_reader.ReadConfigFile(configPath, "user_faker_config.json", &fuc)

	col, err = database.GetCollection("profile")
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "collection not found"))
	}

	fu := user.GenerateUsers(fuc)

	for _, val := range fu {
		err = col.Insert(val)
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "error while adding new user"))
		}
	}
}
