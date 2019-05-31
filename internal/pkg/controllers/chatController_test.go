package controllers

import (
	"net/http"
	"testing"
)

var testsConnectChat = []TestCases{
	// authCtx может быть только тру из-за миддлвара
	{
		routerPath: "/api/chat/global/ws",
		method:     "GET",
		url:        "/api/chat/global/ws",
		body:       nil,
		urlValues:  "",
		expectedRes: `Bad Request
`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
		userIDCtx:      "",
	},
}

func TestConnectToGlobalChat(t *testing.T) {
	MainInit()
	err := testHandler(ConnectToGlobalChat, testsConnectChat, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
