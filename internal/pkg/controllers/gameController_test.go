package controllers

import (
	"net/http"
	"testing"
)

var testsPlay = []TestCases{
	// authCtx может быть только тру из-за миддлвара
	{
		routerPath:     "/api/game/ws",
		method:         "GET",
		url:            "/api/game/ws",
		body:           nil,
		urlValues:      "",
		expectedRes:    `Bad Request
`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
		userIDCtx:      "",
	},
	{
		routerPath:     "/api/game/ws",
		method:         "GET",
		url:            "/api/game/ws",
		body:           nil,
		urlValues:      "",
		expectedRes:    `Bad Request
`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
		userIDCtx:      "",
	},
}

func TestPlay(t *testing.T) {
	//MainInit()
	err := testHandler(Play, testsPlay, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
