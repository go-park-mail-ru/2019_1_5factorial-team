package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsUsersCountInfo = []TestCases{
	{
		routerPath: "/api/user/count",
		method:     "GET",
		url:        "/api/user/count",
		body:       nil,
		urlValues:  "",
		expectedRes:    `{"count":21}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath: "/api/user/count",
		method:     "GET",
		url:        "/api/user/count",
		body:       nil,
		urlValues:  "",
		expectedRes:    `{"count":21}`,
		expectedStatus: http.StatusOK,
		authCtx:        true,
	},
}

func TestUsersCountInfo(t *testing.T) {
	//MainInit()
	err := testHandler(UsersCountInfo, testsUsersCountInfo, t)
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
		expectedRes:    `{"error":"already auth, ctx.authorized shouldn't be true"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Cannot create user: error while adding new user: E11000 duplicate key error collection: user.profile index: nickname_1 dup key: { : \"kekkekkek\" }"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"Cannot create user: error while adding new user: E11000 duplicate key error collection: user.profile index: nickname_1 dup key: { : \"kekkekkek\" }"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek1", "email": "kek@email1.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `"signUp ok"`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
}

func TestSignUp(t *testing.T) {
	err := testHandler(SignUp, testsSignUp, t)
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
		expectedRes:    `{"error":"user with this id not found"}`,
		expectedStatus: http.StatusNotFound,
	},
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/500000000000000000000000000000000000000000000000",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user with this id not found"}`,
		expectedStatus: http.StatusNotFound,
	},
}

func TestGetUser(t *testing.T) {
	//MainInit()
	err := testHandler(GetUser, testsGetUser, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
