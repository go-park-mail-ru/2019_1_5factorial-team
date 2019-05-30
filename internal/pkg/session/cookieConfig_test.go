package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
	"testing"
	"time"
)

func TestUpdateHttpCookie(t *testing.T) {
	c := http.Cookie{}
	exp := time.Now()
	UpdateHttpCookie(&c, exp)

	if c.Path != config.Get().CookieConfig.ServerPrefix {
		t.Error("wrong cookie path")
	}
	if c.Name != config.Get().CookieConfig.CookieName {
		t.Error("wrong cookie name")
	}
	if c.HttpOnly != config.Get().CookieConfig.HttpOnly {
		t.Error("wrong cookie http only")
	}
	if c.Expires != exp {
		t.Error("wrong cookie expiration")
	}
}
