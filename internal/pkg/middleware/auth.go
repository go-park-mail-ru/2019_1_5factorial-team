package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var userId string = ""
		authorized := false

		defer func() {
			ctx = context.WithValue(ctx, "userID", userId)
			ctx = context.WithValue(ctx, "authorized", authorized)

			if authorized {
				ctx = context.WithValue(ctx, "logger", log.LoggerWithAuth(req))
			} else {
				ctx = context.WithValue(ctx, "logger", log.LoggerWithoutAuth(req))
			}

			next.ServeHTTP(res, req.WithContext(ctx))
		}()

		cookie, err := req.Cookie(config.Get().CookieConfig.CookieName)
		if err != nil {
			//log.Println("no cookie found, user unauthorized")
			logrus.WithField("cookie", cookie).Warn("no cookie found, user unauthorized")

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
				// TODO(): переделать на ErrResponse
				http.Error(res, "relogin, please", http.StatusInternalServerError)
			}

			session.UpdateHttpCookie(cookie, updatedToken.CookieExpiredTime)
		}

		http.SetCookie(res, cookie)
	})

}

func CheckLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.Context().Value("logger").(*logrus.Entry).Info("CheckLoginMiddleware")

		// request has context, bcs its coming after AuthMiddleware
		if req.Context().Value("authorized").(bool) == false {
			// TODO(): переделать на ErrResponse
			http.Error(res, "unauthorized, login please", http.StatusUnauthorized)

			logrus.WithField("authorized", req.Context().Value("authorized").(bool)).Warn("user unauthorized")
			return
		}
		next.ServeHTTP(res, req)
	})

}
