package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/oauth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
)

const DefaultPassword = " "

type Token struct {
	Token string `json:"token"`
}

func OauthUser(res http.ResponseWriter, req *http.Request, service string) {

	token := Token{}
	status, err := ParseRequestIntoStruct(true, req, &token)
	if err != nil {
		ErrResponse(res, status, err.Error())

		log.Println("\t", errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}
	client := &http.Client{}
	request, err := http.NewRequest("GET", oauth.GetApiUrl(token.Token, service), nil)

	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println("\t", errors.Wrap(err, "get api url with token "+string(token.Token)+"and service "+string(service)))
		return
	}

	if service == "yandex" {
		request.Header.Set("Authorization", "OAuth "+token.Token)
	}
	response, err := client.Do(request)

	defer response.Body.Close()
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())
		log.Println("\t", errors.Wrap(err, "cant HTTP request"))
		return
	}

	contents, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		log.Println("\t", errors.Wrap(err, "body ReadAll error"))
		return
	}

	uuid, name, err := oauth.GetOauthUser(service, contents)
	if err != nil {
		ErrResponse(res, http.StatusForbidden, err.Error())

		log.Println("\t", errors.Wrap(err, "err in get oauth user"))
		return
	}

	searchingUser, err := user.GetUser(uuid)
	if err != nil {
		searchingUser, err = user.CreateUser(name, uuid, DefaultPassword)
		if err != nil {
			ErrResponse(res, http.StatusBadRequest, err.Error())

			log.Println("\t", errors.Wrap(err, "err in user data"))
			return
		}

	}

	randToken, expiration, err := session.SetToken(searchingUser.Id)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println("\t", errors.Wrap(err, "err set token"))
		return
	}
	cookie := session.CreateHttpCookie(randToken, expiration)

	http.SetCookie(res, cookie)
	OkResponse(res, "oauth ok")
	return

}
