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
		expectedRes:    `{"error":"token is empty"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session/oauth/yandex",
		method:         "POST",
		url:            "/api/session/oauth/yandex",
		body:           strings.NewReader(`{: ""}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character ':' looking for beginning of object key string"}`,
		expectedStatus: http.StatusInternalServerError,
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
