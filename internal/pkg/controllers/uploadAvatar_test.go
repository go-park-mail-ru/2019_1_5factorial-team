package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsUploadAvatar= []TestCases{
	{
		routerPath:     "/api/user",
		method:         "PUT",
		url:            "/api/user",
		body:           strings.NewReader(`{"avatar": "kekkekkek1", "old_password": "1","new_password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"file too big"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
		userIDCtx:      "5556c0d9b49cd4582aaad41c",
	},
}

func TestUploadAvatar(t *testing.T) {
	//MainInit()
	err := testHandler(UploadAvatar, testsUploadAvatar, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
