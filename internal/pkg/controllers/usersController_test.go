package controllers

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type TestCases struct {
	method         string
	url            string
	body           io.Reader
	urlValues      string
	expectedRes    string
	expectedStatus int
	authCtx        bool
}

var testsGetUser = []TestCases{
	{
		method:         "GET",
		url:            "/api/user/",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user id not provided"}`,
		expectedStatus: http.StatusBadRequest,
	},
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
		handler := http.HandlerFunc(funcToTest)

		ctx := req.Context()
		ctx = context.WithValue(ctx, "authorized", val.authCtx)
		req = req.WithContext(ctx)

		handler.ServeHTTP(rr, req)

		// try to check cookie, LOOKS LIKE HACK
		kek := rr.Header()["Set-Cookie"]
		if kek != nil {
			log.Println(strings.Split(kek[0], ";"))
		}

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

var testsUsersCountInfo = []TestCases{
	{
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"count":` + strconv.Itoa(user.GetUsersCount()) + `}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		// сейчас 21, но хз сколько будет
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

// undone
func TestGetUser(t *testing.T) {
	var req *http.Request
	var err error
	for _, val := range testsGetUser {

		req, err = http.NewRequest(val.method, val.url, val.body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetUser)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != val.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, val.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), val.expectedRes)
		}
	}
}

var testsSignUp = []TestCases{
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character '}' looking for beginning of object key string"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Cannot create user: User with this nickname already exist"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
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

// get leaderboard


var testsSignIn = []TestCases{
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Wrong password or login"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"},`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character ',' after top-level value"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"loginOrEmail": "kekkekkek", "password": "kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Wrong password or login"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		method:         "POST",
		url:            "/api/user",
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