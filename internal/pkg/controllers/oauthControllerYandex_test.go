package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsLoginFromYandex = []TestCases{
	{
		routerPath:     "/api/session/oauth/yandex",
		method:         "POST",
		url:            "/api/session/oauth/yandex",
		body:           strings.NewReader(`{"token": "qwqwqwqwqw"}`),
		urlValues:      "",
		expectedRes:    `{"error":"not valid token"}`,
		expectedStatus: http.StatusBadGateway,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session/oauth/yandex",
		method:         "POST",
		url:            "/api/session/oauth/yandex",
		body:           strings.NewReader(`{"token": ""}`),
		urlValues:      "",
		expectedRes:    `{"error":"not valid token"}`,
		expectedStatus: http.StatusBadGateway,
		authCtx:        false,
	},
}

func TestLoginFromYandex(t *testing.T) {
	//MainInit()
	err := testHandler(LoginFromYandex, testsLoginFromYandex, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
