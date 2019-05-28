package controllers

import (
	"net/http"
)

// HelloWorld godoc
// @Title HelloWorld
// @Summary on hello returning world
// @ID hello-world
// @Success 200 {string} World
// @Router /hello [get]
func HelloWorld(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("World"))
	panic("kek")
}
