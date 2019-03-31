package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
	"time"
)

func CreateHttpCookie(value string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Path:     config.Get().CookieConfig.ServerPrefix,
		Name:     config.Get().CookieConfig.CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: config.Get().CookieConfig.HttpOnly,
	}
}

func UpdateHttpCookie(cookie *http.Cookie, expiration time.Time) {
	cookie.Path = config.Get().CookieConfig.ServerPrefix
	cookie.Name = config.Get().CookieConfig.CookieName
	cookie.Expires = expiration
	cookie.HttpOnly = config.Get().CookieConfig.HttpOnly
}
