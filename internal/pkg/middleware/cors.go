package middleware

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CORSConfig struct {
	Origin      string   `json:"allow-origin"`
	Credentials bool     `json:"allow-credentials"`
	Methods     []string `json:"allow-methods"`
	Headers     []string `json:"allow-headers"`
	MaxAge      int      `json:"max-age"`
}

var config = CORSConfig{}

func init() {
	err := config_reader.ReadConfigFile("cors_config.json", &config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading CORS config"))
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL, "CORS")

		val, ok := req.Header["Origin"]
		if ok {
			if config.Origin == "*" {
				res.Header().Set("Access-Control-Allow-Origin", val[0]) // "*"
			} else {
				res.Header().Set("Access-Control-Allow-Origin", config.Origin)
			}
			res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.Credentials))
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(config.Methods, ","))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(config.Headers, ","))
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.MaxAge)) // 24 hours
			return
		}

		next.ServeHTTP(res, req)
	})
}
