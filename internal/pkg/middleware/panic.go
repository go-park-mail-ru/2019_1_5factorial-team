package middleware

import (
	"log"
	"net/http"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("OOOOOPA PANIC recovered", err)
				http.Error(res, "some error happened", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(res, req)
	})
}
