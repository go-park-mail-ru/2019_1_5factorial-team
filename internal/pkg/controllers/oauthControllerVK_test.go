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
		expectedRes:    `{"error":"token is empty"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session/oauth/vk",
		method:         "POST",
		url:            "/api/session/oauth/vk",
		body:           strings.NewReader(`{: ""}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character ':' looking for beginning of object key string"}`,
		expectedStatus: http.StatusInternalServerError,
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
