package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsLoginFromGoogle = []TestCases{
	{
		routerPath:     "/api/session/oauth/google",
		method:         "POST",
		url:            "/api/session/oauth/google",
		body:           strings.NewReader(`{"token": "qwqwqwqwqw"}`),
		urlValues:      "",
		expectedRes:    `"oauth ok"`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
}

func TestLoginFromGoogle(t *testing.T) {
	//MainInit()
	err := testHandler(LoginFromGoogle, testsLoginFromGoogle, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
