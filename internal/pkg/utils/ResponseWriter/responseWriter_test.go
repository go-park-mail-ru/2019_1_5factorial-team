package ResponseWriter

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusWriter_WriteHeader(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		nRes := NewStatusWriter(res)
		nRes.WriteHeader(http.StatusBadRequest)

		if nRes.status != http.StatusBadRequest {
			t.Error("wrong status in statusWriter")
		}
	})

	router := mux.NewRouter()
	router.HandleFunc("/", handler)

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(httptest.NewRecorder(), req)
}
