package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctxLogger := req.Context().Value("logger").(*logrus.Entry)
		if ctxLogger == nil {
			t.Error("ctxLogger is nill")
			return
		}
	})

	middleware := AuthMiddleware(handler)
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	middleware.ServeHTTP(httptest.NewRecorder(), req)
}
