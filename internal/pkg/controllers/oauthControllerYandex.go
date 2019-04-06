package controllers

import (
	"net/http"
)

func LoginFromYandex(res http.ResponseWriter, req *http.Request) {
	OauthUser(res, req, "yandex")
}
