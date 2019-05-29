package middleware

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitConfig() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	log.InitLogs()
}

func TestCORSMiddleware(t *testing.T) {
	InitConfig()
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t.Error("this shouldnt happen")
	})

	middleware := CORSMiddleware(handler)
	// create a mock request to use
	req := httptest.NewRequest("OPTIONS", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	res := httptest.NewRecorder()
	middleware.ServeHTTP(res, req)
}
