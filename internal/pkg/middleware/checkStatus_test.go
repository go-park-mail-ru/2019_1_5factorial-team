package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatus(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	})

	middleware := CheckStatus(handler)
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	middleware.ServeHTTP(httptest.NewRecorder(), req)
}
