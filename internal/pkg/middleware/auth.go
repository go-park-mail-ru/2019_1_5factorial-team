package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		cookie, err := req.Cookie("token")
		if err != nil {

			ctx = context.WithValue(ctx, "userID", -1)
			ctx = context.WithValue(ctx, "authorized", false)
			ctx = context.WithValue(ctx, "token", "")
		} else {

			uId, err := session.GetId(cookie.Value)
			// TODO(smet1):
			//if err != nil {
			//	// invalid token
			//	// KILL HIM
			//	// TODO(smet1): ничего умнее не придумал
			//
			//	cookie.Expires = time.Unix(0, 0)
			//	http.SetCookie(res, cookie)
			//	http.Error(res, "cookie invalid, relogin please", http.StatusTeapot)
			//	return
			//}

			if err != nil {
				ctx = context.WithValue(ctx, "userID", -1)
				ctx = context.WithValue(ctx, "authorized", false)
				ctx = context.WithValue(ctx, "token", "")

				cookie.Expires = time.Unix(0, 0)
				http.SetCookie(res, cookie)

			} else {
				ctx = context.WithValue(ctx, "userID", uId)
				ctx = context.WithValue(ctx, "authorized", true)
				ctx = context.WithValue(ctx, "token", cookie.Value)

				// сетим новое время куки
				// юзер может придти, когда его токен испортиться в момент обращения к серверу,
				// поэтому добавлю его текущий токен
				updatedToken, err := session.UpdateToken(cookie.Value, uId)
				if err != nil {
					http.Error(res, "relogin, please", http.StatusInternalServerError)
				}
				cookie.Expires = updatedToken.CookieExpiredTime
				http.SetCookie(res, cookie)
			}
		}

		next.ServeHTTP(res, req.WithContext(ctx))
	})

}
