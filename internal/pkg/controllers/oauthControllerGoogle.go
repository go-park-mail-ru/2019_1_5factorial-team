package controllers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/oauth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/pkg/errors"
)

func LoginFromGoogle(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)

	token := oauth.Token{}
	status, err := ParseRequestIntoStruct(true, req, &token)
	if err != nil {
		ErrResponse(res, status, err.Error())

		ctxLogger.Println("\t", errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}

	status, err, randToken, expiration := oauth.OauthUser(token.Token, "google")
	if err != nil {
		ErrResponse(res, status, err.Error())

		ctxLogger.Println("\t", err)
		return
	}

	cookie := session.CreateHttpCookie(randToken, expiration)

	http.SetCookie(res, cookie)
	OkResponse(res, nil)
}
