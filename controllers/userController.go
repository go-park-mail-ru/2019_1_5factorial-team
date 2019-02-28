package controllers

import (
	"fmt"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}
