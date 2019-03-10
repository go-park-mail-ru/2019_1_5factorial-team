package controllers

import (
	"fmt"
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

func GetLeaderboard(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	fmt.Println(query)
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	fmt.Println(limit, offset)

	leaderboard, err := user.GetUsersScores(limit, offset)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println(errors.Wrap(err, "get leaderboard error"))
		return
	}

	fmt.Println(leaderboard)
	OkResponse(res, GetLeaderboardResponse{
		Scores: leaderboard,
	})
}
