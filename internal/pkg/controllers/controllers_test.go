package controllers

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/gorilla/mux"
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
	TestGetUserFromSession,
	TestGetLeaderboard,
	TestUsersCountInfo,
	TestGetUser,
	TestSignUp,
	TestSignIn,
}

func TestControllers(t *testing.T) {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		t.Error(err.Error())
	}

	log.InitLogs()

	database.InitConnection()

	for _, test := range funcs {
		test(t)
	}
}

func testHandler(funcToTest func(http.ResponseWriter, *http.Request), tests []TestCases, t *testing.T) error {
	var req *http.Request
	var err error
	for _, val := range tests {

		req, err = http.NewRequest(val.method, val.url, val.body)
		if err != nil {
			return err
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc(val.routerPath, funcToTest).Methods(val.method)

		ctx := req.Context()
		ctx = context.WithValue(ctx, "authorized", val.authCtx)
		ctx = context.WithValue(ctx, "userID", val.userIDCtx)
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
			t.Errorf("handler returned wrong status code: \ngot  %v \nwant %v",
				status, val.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			t.Errorf("handler returned unexpected body: \ngot  %v \nwant %v",
				rr.Body.String(), val.expectedRes)
		}
	}
	return nil
}
