package controllers

import "net/http"

func HW(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("World"))
}
