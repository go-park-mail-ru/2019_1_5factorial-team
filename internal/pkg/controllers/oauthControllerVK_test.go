package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsLoginFromVK = []TestCases{
	{
		routerPath:     "/api/session/oauth/vk",
		method:         "POST",
		url:            "/api/session/oauth/vk",
		body:           strings.NewReader(`{"token": "qwqwqwqwqw"}`),
		urlValues:      "",
		expectedRes:    `{"error":"err in get oauth user: not valid token"}`,
		expectedStatus: http.StatusForbidden,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session/oauth/vk",
		method:         "POST",
		url:            "/api/session/oauth/vk",
		body:           strings.NewReader(`{"token": ""}`),
		urlValues:      "",
		expectedRes:    `{"error":"err in get oauth user: not valid token"}`,
		expectedStatus: http.StatusForbidden,
		authCtx:        false,
	},
}

func TestLoginFromVK(t *testing.T) {
	//MainInit()
	err := testHandler(LoginFromVK, testsLoginFromVK, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
