package controllers

import (
	"net/http"
)

func LoginFromYandex(res http.ResponseWriter, req *http.Request) {
	GetUserInfo(res, req, "yandex")
}
