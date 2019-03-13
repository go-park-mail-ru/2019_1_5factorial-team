package session

import (
	"net/http"
	"time"
)

const (
	CookieName      string = "token"
	HttpOnly        bool   = true
	CookieTimeHours        = 10
	ServerPrefix    string = "/api"
)

func CreateHttpCookie(value string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Path:     ServerPrefix,
		Name:     CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: HttpOnly,
	}
}

func UpdateHttpCookie(cookie *http.Cookie, expiration time.Time) {
	cookie.Path = ServerPrefix
	cookie.Name = CookieName
	cookie.Expires = expiration
	cookie.HttpOnly = HttpOnly
}
