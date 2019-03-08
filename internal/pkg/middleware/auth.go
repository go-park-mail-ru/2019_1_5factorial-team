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

		cookie, err := req.Cookie(session.CookieName)
		if err != nil {

			ctx = context.WithValue(ctx, "userID", -1)
			ctx = context.WithValue(ctx, "authorized", false)
			ctx = context.WithValue(ctx, "token", "")
		} else {

			uId, err := session.GetId(cookie.Value)

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
				// и обновляем время токена
				updatedToken, err := session.UpdateToken(cookie.Value)
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

func CheckLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// request has context, bcs its coming after AuthMiddleware
		if req.Context().Value("authorized").(bool) == false {
			http.Error(res, "unauthorized, login please", http.StatusUnauthorized)

			return
		}
		next.ServeHTTP(res, req)
	})

}
