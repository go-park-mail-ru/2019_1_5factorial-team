package middleware

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//type CORSConfig struct {
//	Origin      string   `json:"allow-origin"`
//	Credentials bool     `json:"allow-credentials"`
//	Methods     []string `json:"allow-methods"`
//	Headers     []string `json:"allow-headers"`
//	MaxAge      int      `json:"max-age"`
//}

//var CorsConfig = CORSConfig{}

//func init() {
//	err := config_reader.ReadConfigFile("cors_config.json", &CorsConfig)
//	if err != nil {
//		log.Fatal(errors.Wrap(err, "error while reading CORS config"))
//	}
//}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL, "CORS")

		val, ok := req.Header["Origin"]
		if ok {
			if config.GetInstance().CORSConfig.Origin == "*" {
				res.Header().Set("Access-Control-Allow-Origin", val[0]) // "*"
			} else {
				res.Header().Set("Access-Control-Allow-Origin", config.GetInstance().CORSConfig.Origin)
			}
			res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.GetInstance().CORSConfig.Credentials))
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(config.GetInstance().CORSConfig.Methods, ","))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(config.GetInstance().CORSConfig.Headers, ","))
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.GetInstance().CORSConfig.MaxAge)) // 24 hours

			log.Println("OPTIONS")
			return
		}

		next.ServeHTTP(res, req)
	})
}
