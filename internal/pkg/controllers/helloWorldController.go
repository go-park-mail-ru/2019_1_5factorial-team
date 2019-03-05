package controllers

import (
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("World"))
}
