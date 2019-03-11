package session

import (
	"net/http"
	"time"
)

const (
	CookieName      string = "token"
	HttpOnly        bool   = true
	CookieTimeHours        = 10
)

func CreateHttpCookie(value string, expiration time.Time) http.Cookie {
	return http.Cookie{
		Name:     CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: HttpOnly,
	}
}
