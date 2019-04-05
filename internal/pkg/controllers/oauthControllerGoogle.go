package controllers

import (
	"net/http"
)

func LoginFromGoogle(res http.ResponseWriter, req *http.Request) {
	GetUserInfo(res, req, "google")
}
