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
		expectedRes:    `{"error":"json parsing error: parse error: invalid character ',' after top-level value near offset 33 of ','"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session",
		method:         "POST",
		url:            "/api/session",
		body:           strings.NewReader(`{"loginOrEmail": "kek@email.kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth, ctx.authorized shouldn't be true"}`,
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
		expectedRes:    ``,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
}

func TestSignIn(t *testing.T) {
	//MainInit()
	err := testHandler(SignIn, testsSignIn, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsGetUserFromSession = []TestCases{
	// authCtx может быть только тру из-за миддлвара
	{
		routerPath:     "/api/user",
		method:         "GET",
		url:            "/api/user",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"invalid empty id"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
		userIDCtx:      "",
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
		userIDCtx:      "-1",
	},
}

func TestGetUserFromSession(t *testing.T) {
	//MainInit()
	err := testHandler(GetUserFromSession, testsGetUserFromSession, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsSignOut = []TestCases{
	// authCtx может быть только тру из-за миддлвара
	{
		routerPath:     "/api/session",
		method:         "DELETE",
		url:            "/api/session",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"not authorized"}`,
		expectedStatus: http.StatusUnauthorized,
		authCtx:        true,
		userIDCtx:      "",
	},
	{
		routerPath:     "/api/session",
		method:         "DELETE",
		url:            "/api/session",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"not authorized"}`,
		expectedStatus: http.StatusUnauthorized,
		authCtx:        false,
		userIDCtx:      "",
	},
}

func TestSignOut(t *testing.T) {
	//MainInit()
	err := testHandler(SignOut, testsSignOut, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsIsSessionValid = []TestCases{
	// authCtx может быть только тру из-за миддлвара
	{
		routerPath:     "/api/session",
		method:         "GET",
		url:            "/api/session",
		body:           nil,
		urlValues:      "",
		expectedRes:    ``,
		expectedStatus: http.StatusOK,
		authCtx:        true,
		userIDCtx:      "",
	},
}

func TestIsSessionValid(t *testing.T) {
	//MainInit()
	err := testHandler(IsSessionValid, testsIsSessionValid, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
