package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"log"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL, "AuthMiddleware")

		ctx := req.Context()
		var userId int64 = -1
		authorized := false

		defer func() {
			ctx = context.WithValue(ctx, "userID", userId)
			ctx = context.WithValue(ctx, "authorized", authorized)

			next.ServeHTTP(res, req.WithContext(ctx))
		}()

		cookie, err := req.Cookie(session.CookieConf.CookieName)
		if err != nil {
			log.Println("no cookie found, user unauthorized")
			return
		}

		uId, err := session.GetId(cookie.Value)

		if err != nil {
			cookie.Expires = time.Unix(0, 0)

		} else {
			userId = uId
			authorized = true

			// сетим новое время куки
			// и обновляем время токена
			updatedToken, err := session.UpdateToken(cookie.Value)
			if err != nil {
				http.Error(res, "relogin, please", http.StatusInternalServerError)
			}

			session.UpdateHttpCookie(cookie, updatedToken.CookieExpiredTime)
		}

		http.SetCookie(res, cookie)
	})

}

func CheckLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL, "CheckLoginMiddleware")
		// request has context, bcs its coming after AuthMiddleware
		if req.Context().Value("authorized").(bool) == false {
			http.Error(res, "unauthorized, login please", http.StatusUnauthorized)

			return
		}
		next.ServeHTTP(res, req)
	})

}
