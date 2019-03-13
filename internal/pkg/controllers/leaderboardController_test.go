package controllers

import (
	"net/http"
	"testing"
)

var testsGetLeaderboard = []TestCases{
	{
		routerPath:     "/api/user/score",
		method:         "POST",
		url:            "/api/user/score?limit=1&offset=1",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"scores":[{"nickname":"kekkekkek","score":100500}]}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user/score",
		method:         "POST",
		url:            "/api/user/score?limit=1000&offset=1000",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"limit * (offset - 1) \u003e users count"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
}

func TestGetLeaderboard(t *testing.T) {
	err := testHandler(GetLeaderboard, testsGetLeaderboard, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
