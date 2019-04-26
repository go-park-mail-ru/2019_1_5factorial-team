package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
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
				ctx = context.WithValue(ctx, "logger", log.LoggerWithAuth(req.WithContext(ctx)))
			} else {
				ctx = context.WithValue(ctx, "logger", log.LoggerWithoutAuth(req.WithContext(ctx)))
			}

			next.ServeHTTP(res, req.WithContext(ctx))
		}()

		cookie, err := req.Cookie(config.Get().CookieConfig.CookieName)
		if err != nil {
			//log.Println("no cookie found, user unauthorized")
			logrus.WithField("cookie", cookie).Warn("no cookie found, user unauthorized")

			return
		}

		//uId, err := session.GetId(cookie.Value)
		grpcConn, err := grpcAuth.CreateConnection()
		if err != nil {
			http.Error(res, "relogin, please", http.StatusInternalServerError)

			logrus.Error(errors.Wrap(err, "cant get connection to auth service"))
			return
		}
		defer grpcConn.Close()

		AuthGRPC := grpcAuth.NewAuthCheckerClient(grpcConn)
		ctxGRPC := context.Background()
		uId, err := AuthGRPC.GetIDFromSession(ctxGRPC, &grpcAuth.Cookie{Token: cookie.Value})

		if err != nil {
			cookie.Expires = time.Unix(0, 0)

		} else {
			userId = uId.ID
			authorized = true

			// сетим новое время куки
			// и обновляем время токена
			//updatedToken, err := session.UpdateToken(cookie.Value)
			//grpcConn, err := grpcAuth.CreateConnection()
			//if err != nil {
			//	http.Error(res, "relogin, please", http.StatusInternalServerError)
			//
			//	logrus.Error(errors.Wrap(err, "cant get connection to auth service"))
			//	return
			//}
			//defer grpcConn.Close()
			//
			//AuthGRPC := grpcAuth.NewAuthCheckerClient(grpcConn)
			//ctx := context.Background()
			cookieGRPC, err := AuthGRPC.UpdateSession(ctx, &grpcAuth.Cookie{Token: cookie.Value})
			if err != nil {
				http.Error(res, "relogin, please", http.StatusInternalServerError)

				logrus.Error(errors.Wrap(err, "Set token from grpc returned error"))
				return
			}

			timeCookie, err := time.Parse(time.RFC3339, cookieGRPC.Expiration)
			if err != nil {
				http.Error(res, "relogin, please", http.StatusInternalServerError)

				logrus.Error(errors.Wrap(err, "cant convert time from string"))
				return
			}
			//if err != nil {
			//	TODO(): переделать на ErrResponse
			//http.Error(res, "relogin, please", http.StatusInternalServerError)
			//}

			session.UpdateHttpCookie(cookie, timeCookie)
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
