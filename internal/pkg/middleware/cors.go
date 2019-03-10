package middleware

import (
	"log"
	"net/http"
	"strconv"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL, "CORS")

		val, ok := req.Header["Origin"]
		if ok {
			res.Header().Set("Access-Control-Allow-Origin", val[0])  // "*"
			res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(true))
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
			res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(86400))  // 24 hours
			return
		}

		next.ServeHTTP(res, req)
	})
}
