package controllers

import (
	"net/http"
	"strings"
	"testing"
)

// https://velox-qr.herokuapp.com/login/google/#state=state_parameter_passthrough_value&access_token=ya29.Gl0DBwUoEVa4y4tSLjOb5FLJKDzuSS_a8nzvnEsU8S_4L0OfgnDN4pzQyHmmZRkd8oZr6JGbY1EBkZnlNV24H7p1QSJbIG501DoKMzDy99IwY33oG60fEx1URseQqCQ&token_type=Bearer&expires_in=3600&scope=email%20https://www.googleapis.com/auth/userinfo.email%20openid&authuser=0&session_state=4584ae30923fcbf152207a0daf08e8d6b0cc4863..57f7&prompt=none
var testsLoginFromGoogle = []TestCases{
	{
		routerPath:     "/api/session/oauth/google",
		method:         "POST",
		url:            "/api/session/oauth/google",
		body:           strings.NewReader(`{"token": ""}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: invalid character '-' in numeric literal"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		routerPath:     "/api/session/oauth/google",
		method:         "POST",
		url:            "/api/session/oauth/google",
		body:           strings.NewReader(`{"token": "ya29.Gl0DBwUoEVa4y4tSLjOb5FLJKDzuSS_a8nzvnEsU8S_4L0OfgnDN4pzQyHmmZRkd8oZr6JGbY1EBkZnlNV24H7p1QSJbIG501DoKMzDy99IwY33oG60fEx1URseQqCQ"}`),
		urlValues:      "",
		expectedRes:    `{"count":21}`,
		expectedStatus: http.StatusInternalServerError,
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
