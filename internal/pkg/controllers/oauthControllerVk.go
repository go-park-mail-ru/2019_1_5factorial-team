package controllers

import (
	"net/http"
)

func LoginFromVK(res http.ResponseWriter, req *http.Request) {
	OauthUser(res, req, "vk")
}
