package controllers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCases struct {
	method         string
	url            string
	body           io.Reader
	urlValues      string
	expectedRes    string
	expectedStatus int
	authCtx        bool
}

var testsGetUser = []TestCases{
	{
		method:         "GET",
		url:            "/api/user/",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user id not provided"}`,
		expectedStatus: http.StatusBadRequest,
	},
}

func TestGetUser(t *testing.T) {
	var req *http.Request
	var err error
	for _, val := range testsGetUser {

		req, err = http.NewRequest(val.method, val.url, val.body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetUser)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != val.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, val.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), val.expectedRes)
		}
	}
}

var testsSignUp = []TestCases{
	{
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{{"error":"json parsing error: invalid character '}' looking for beginning of object key string"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
}

func TestSignUp(t *testing.T) {
	var req *http.Request
	var err error
	for _, val := range testsGetUser {

		req, err = http.NewRequest(val.method, val.url, val.body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SignUp)

		ctx := req.Context()
		ctx = context.WithValue(ctx, "authorized", val.authCtx)
		req = req.WithContext(ctx)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != val.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, val.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != val.expectedRes {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), val.expectedRes)
		}
	}
}
