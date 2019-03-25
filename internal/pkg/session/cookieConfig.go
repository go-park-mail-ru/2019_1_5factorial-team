package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

const (
	CookieName      string = "token"
	HttpOnly        bool   = true
	CookieTimeHours        = 10
	ServerPrefix    string = "/api"
)

type CookieConfig struct {
	CookieName      string `json:"cookie_name"`
	HttpOnly        bool   `json:"http_only"`
	CookieDuration  int64  `json:"cookie_time_hours"`
	ServerPrefix    string `json:"server_prefix"`
	CookieTimeHours time.Duration
}

var CookieConf = CookieConfig{}

func init() {
	err := config_reader.ReadConfigFile("cookie_config.json", &CookieConf)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	}
	fmt.Println(CookieConf)

	// чтобы не копипастить везде (это время жизни куки в часах, чтобы добавлять к уже созданной)
	CookieConf.CookieTimeHours = time.Duration(CookieConf.CookieDuration * int64(time.Hour))
}

func CreateHttpCookie(value string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Path:     CookieConf.ServerPrefix,
		Name:     CookieConf.CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: CookieConf.HttpOnly,
	}
}

func UpdateHttpCookie(cookie *http.Cookie, expiration time.Time) {
	cookie.Path = CookieConf.ServerPrefix
	cookie.Name = CookieConf.CookieName
	cookie.Expires = expiration
	cookie.HttpOnly = CookieConf.HttpOnly
}
