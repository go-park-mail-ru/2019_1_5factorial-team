package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
)

const DefaultPassword = ""

type Token struct {
	Token string `json:"token"`
}

func GetUserInfo(res http.ResponseWriter, req *http.Request, service string) error {

	token := Token{}
	status, err := ParseRequestIntoStruct(true, req, &token)
	if err != nil {
		ErrResponse(res, status, err.Error())
		log.Println("\t", errors.Wrap(err, "ParseRequestIntoStruct error"))
		return err
	}
	client := &http.Client{}
	request, _ := http.NewRequest("GET", GetApiUrl(token.Token, service), nil)
	if service == "yandex" {
		request.Header.Set("Authorization", "OAuth "+token.Token)
	}
	response, err := client.Do(request)
	defer response.Body.Close()

	if err != nil {
		ErrResponse(res, http.StatusForbidden, "invalid token")
		return fmt.Errorf("failed getting user info: %s", err.Error())
	}

	contents, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || err != nil {
		ErrResponse(res, response.StatusCode, "invalid token")
		return err
	}

	uuid, name, err := GetOauthUser(service, contents)
	if err != nil {
		ErrResponse(res, http.StatusForbidden, "invalid token")
		return fmt.Errorf("failed reading response body: %s", err.Error())
	}

	searchingUser, err := user.GetUser(uuid)
	if err != nil {
		searchingUser, err = user.CreateUser(name, uuid, DefaultPassword)
	}

	randToken, expiration, err := session.SetToken(searchingUser.Id)

	cookie := session.CreateHttpCookie(randToken, expiration)

	http.SetCookie(res, cookie)
	OkResponse(res, "oauth ok")
	return nil

}
