package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
	ctxLogger := log.WithFields(log.Fields{
		"req":    req.URL,
		"method": req.Method,
		"host":   req.Host,
		"func":   "GetLeaderboard",
	})
	ctxLogger.Info("============================================")

	query := req.URL.Query()

	ctxLogger.Info("query = ", query)

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	leaderboard, err := user.GetUsersScores(limit, offset)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		ctxLogger.Error(errors.Wrap(err, "get leaderboard error"))
		return
	}

	OkResponse(res, GetLeaderboardResponse{
		Scores: leaderboard,
	})

	ctxLogger.Info("OK response = ", leaderboard)
}
