package controllers

import (
	"io"
	"net/http"
	"testing"
)

type TestCasesResp struct {
	routerPath     string
	method         string
	url            string
	body           io.Reader
	urlValues      string
	expectedRes    string
	expectedStatus int
	authCtx        bool
	userIDCtx      string
}

var testsOkResponse = []TestCasesResp{
	{
		routerPath:     "/api/user",
		method:         "GET",
		url:            "/api/user",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"invalid empty id"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
		userIDCtx:      "",
	},
	{
		routerPath:     "/api/user",
		method:         "GET",
		url:            "/api/user",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"not authorized: http: named cookie not present"}`,
		expectedStatus: http.StatusUnauthorized,
		authCtx:        true,
		userIDCtx:      "-1",
	},
}

func TestOkResponse(t *testing.T) {
	for _, val := range testsOkResponse {
		res := http.ResponseWriter()
		OkResponse()
	}
}
