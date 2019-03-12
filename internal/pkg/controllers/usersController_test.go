package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCases struct {
	method         string
	url            string
	body           io.Reader
	urlValues      string
	expectedRes    string
	expectedStatus int
}

var test = []TestCases{
	{
		method:         "GET",
		url:            "/api/user/",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user id not provided"}`,
		expectedStatus: http.StatusBadRequest,
	},
}

func TestSignUp(t *testing.T) {
	var req *http.Request
	var err error
	for _, val := range test {

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
