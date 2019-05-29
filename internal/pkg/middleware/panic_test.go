package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanicMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		sl := make([]int, 0, 0)
		sl[100500] = 1
	})

	middleware := PanicMiddleware(handler)
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	res := httptest.NewRecorder()
	middleware.ServeHTTP(res, req)
}
