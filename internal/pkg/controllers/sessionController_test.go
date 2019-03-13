package controllers

import (
	"net/http"
	"strings"
	"testing"
)

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
	err := testHandler(SignIn, testsSignIn, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsGetUserFromSession = []TestCases{
	{
		routerPath:     "/api/user",
		method:         "GET",
		url:            "/api/user",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"email":"kek.k.ek","nickname":"kekkekkek","score":100500,"avatar_link":"../../../img/default.jpg"}`,
		expectedStatus: http.StatusOK,
		authCtx:        true,
		userIDCtx:      0,
	},
	{
		routerPath:     "/api/user",
		method:         "GET",
		url:            "/api/user",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"not authorized: http: named cookie not present"}`,
		expectedStatus: http.StatusUnauthorized,
		authCtx:        true,
		userIDCtx:      -1,
	},
}

func TestGetUserFromSession(t *testing.T) {
	err := testHandler(GetUserFromSession, testsGetUserFromSession, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
