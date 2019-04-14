package middleware

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
	"strconv"
	"strings"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//log.Println(req.URL, "CORS")

		val, ok := req.Header["Origin"]
		if ok {
			if config.Get().CORSConfig.Origin == "*" {
				res.Header().Set("Access-Control-Allow-Origin", val[0]) // "*"
			} else {
				res.Header().Set("Access-Control-Allow-Origin", config.Get().CORSConfig.Origin)
			}
			res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.Get().CORSConfig.Credentials))
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(config.Get().CORSConfig.Methods, ","))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(config.Get().CORSConfig.Headers, ","))
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.Get().CORSConfig.MaxAge)) // 24 hours

			//log.Println("OPTIONS")
			return
		}

		next.ServeHTTP(res, req)
	})
}
