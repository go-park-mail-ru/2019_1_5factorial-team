package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithField("err", err).Error("OOOOOPA PANIC MIDDLEWARE, recovered")
				http.Error(res, "some error happened", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(res, req)
	})
}
