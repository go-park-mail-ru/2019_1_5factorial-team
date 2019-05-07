package controllers

import (
	"context"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

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
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)
	ctx := context.Background()

	query := req.URL.Query()

	ctxLogger.Info("query = ", query)

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	scores, err := authGRPC.GetUsersScores(ctx, &grpcAuth.ScoresParam{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		ctxLogger.Error(errors.Wrap(err, "get leaderboard error"))
		return
	}

	leaderboard := make([]user.Scores, 0, len(scores.Scores))
	for _, val := range scores.Scores {
		leaderboard = append(leaderboard, user.Scores{Nickname: val.Nickname, Score: int(val.Score)})
	}

	OkResponse(res, GetLeaderboardResponse{
		Scores: leaderboard,
	})

	ctxLogger.Info("OK response = ", leaderboard)
}
