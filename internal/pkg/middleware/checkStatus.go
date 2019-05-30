package middleware

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/stats"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/ResponseWriter"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"net/http"
)

func CheckStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		stats.Stats.AddHit()

		myRW := ResponseWriter.NewStatusWriter(res)
		next.ServeHTTP(myRW, req)

		log.Println(myRW.GetStatusCode())
		if myRW.GetStatusCode() != http.StatusOK {
			stats.Stats.AddBadResponse(myRW.GetStatusCode(), req.URL.Path)
		}
	})
}
