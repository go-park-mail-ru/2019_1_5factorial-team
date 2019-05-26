package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/stats"
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
	stats.Hits.WithLabelValues("200", req.URL.String()).Inc()
	panic("kek")
}
