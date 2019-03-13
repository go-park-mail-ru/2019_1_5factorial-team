package controllers

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
}

func testHandler(funcToTest func(http.ResponseWriter, *http.Request), tests []TestCases) error {
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
		req = req.WithContext(ctx)

		router.ServeHTTP(rr, req)

		// try to check cookie, LOOKS LIKE HACK
		//kek := rr.Header()["Set-Cookie"]
		//if kek != nil {
		//	log.Println(strings.Split(kek[0], ";"))
		//}

		// Check the status code is what we expect.
		if status := rr.Code; status != val.expectedStatus {
			return errors.Errorf("handler returned wrong status code: \ngot  %v \nwant %v",
				status, val.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			return errors.Errorf("handler returned unexpected body: \ngot  %v \nwant %v",
				rr.Body.String(), val.expectedRes)
		}
	}
	return nil
}

var testsGetLeaderboard = []TestCases{
	{
		routerPath:     "/api/user/score",
		method:         "POST",
		url:            "/api/user/score?limit=1&offset=1",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"scores":[{"nickname":"kekkekkek","score":100500}]}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user/score",
		method:         "POST",
		url:            "/api/user/score?limit=1000&offset=1000",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"limit * (offset - 1) \u003e users count"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
}

func TestGetLeaderboard(t *testing.T) {
	err := testHandler(GetLeaderboard, testsGetLeaderboard)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsUsersCountInfo = []TestCases{
	{
		routerPath:     "/api/user/count",
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"count":` + strconv.Itoa(user.GetUsersCount()) + `}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user/count",
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"count":` + strconv.Itoa(user.GetUsersCount()) + `}`,
		expectedStatus: http.StatusOK,
		authCtx:        true,
	},
}

func TestUsersCountInfo(t *testing.T) {
	err := testHandler(UsersCountInfo, testsUsersCountInfo)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsGetUser = []TestCases{
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/0",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"email":"kek.k.ek","nickname":"kekkekkek","score":100500,"avatar_link":"../../../img/default.jpg"}`,
		expectedStatus: http.StatusOK,
	},
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/50",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user with this id not found"}`,
		expectedStatus: http.StatusNotFound,
	},
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/500000000000000000000000000000000000000000000000",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"bad id"}`,
		expectedStatus: http.StatusInternalServerError,
	},
}

func TestGetUser(t *testing.T) {
	err := testHandler(GetUser, testsGetUser)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsSignUp = []TestCases{
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character '}' looking for beginning of object key string"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Cannot create user: User with this nickname already exist"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek1", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `"signUp ok"`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
}

func TestSignUp(t *testing.T) {
	err := testHandler(SignUp, testsSignUp)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsSignIn = []TestCases{
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Wrong password or login"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"},`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character ',' after top-level value"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kekkekkek", "password": "kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Wrong password or login"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kekkekkek", "password": "password"}`),
		urlValues:      "",
		expectedRes:    `"ok auth"`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
}

func TestSignIn(t *testing.T) {
	err := testHandler(SignIn, testsSignIn)
	if err != nil {
		t.Errorf(err.Error())
	}
}
