package controllers

import (
	"net/http"
)

func LoginFromVK(res http.ResponseWriter, req *http.Request) {
	GetUserInfo(res, req, "vk")
}
