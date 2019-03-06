package server

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"net/http"
	"time"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		cookie, err := req.Cookie("token")
		if err != nil {

			ctx = context.WithValue(ctx,"realId", -1)
			ctx = context.WithValue(ctx, "authorized", false)
		} else {

			uId, err := session.GetId(cookie.Value)
			if err != nil {
				// invalid token
				// KILL HIM
				// TODO(smet1): ничего умнее не придумал

				cookie.Expires = time.Now().AddDate(0, 0, -1)
				http.SetCookie(res, cookie)
				http.Error(res, "cookie invalid", http.StatusTeapot)
				return
			}

			ctx = context.WithValue(ctx,"realId", uId)
			ctx = context.WithValue(ctx, "authorized", true)
			ctx = context.WithValue(ctx, "token", cookie.Value)

			// сетим новое время куки
			cookie.Expires = time.Now().Add(10 * time.Hour)
			http.SetCookie(res, cookie)
		}

		next.ServeHTTP(res, req.WithContext(ctx))
	})
	
}
