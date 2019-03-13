package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"email":
// 	"nickname":
// 	"score":
// 	"avatar_link":

type GetLeaderboardResponse struct {
	Scores []user.Scores `json:"scores"`
}

// GetLeaderboard godoc
// @Title get leader board
// @Summary return slice of Scores (Nickname + score)
// @ID get-leaderboard
// @Produce json
// @Param offset query int false "default: 0"
// @Param limit query int false "default: 10"
// @Success 200 {array} controllers.GetLeaderboardResponse
// @Failure 400 {object} controllers.errorResponse
// @Router /user/score [get]
func GetLeaderboard(res http.ResponseWriter, req *http.Request) {
	log.Println("================", req.URL, req.Method, "GetLeaderboard", "================")

	query := req.URL.Query()
	log.Println("\t query = ", query)
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	leaderboard, err := user.GetUsersScores(limit, offset)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println("\t", errors.Wrap(err, "get leaderboard error"))
		return
	}

	OkResponse(res, GetLeaderboardResponse{
		Scores: leaderboard,
	})
	log.Println("\t", "ok response GetLeaderboard")
	for i, val := range leaderboard {
		log.Printf("\t\t i = %d, nickname = %s, score = %d", i, val.Nickname, val.Score)
	}
}
