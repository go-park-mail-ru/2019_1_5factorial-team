package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
	"time"
)

func CreateHttpCookie(value string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Path:     config.GetInstance().CookieConfig.ServerPrefix,
		Name:     config.GetInstance().CookieConfig.CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: config.GetInstance().CookieConfig.HttpOnly,
	}
}

func UpdateHttpCookie(cookie *http.Cookie, expiration time.Time) {
	cookie.Path = config.GetInstance().CookieConfig.ServerPrefix
	cookie.Name = config.GetInstance().CookieConfig.CookieName
	cookie.Expires = expiration
	cookie.HttpOnly = config.GetInstance().CookieConfig.HttpOnly
}
