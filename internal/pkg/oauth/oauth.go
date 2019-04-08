package oauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
)

const DefaultPassword = " "

type Token struct {
	Token string `json:"token"`
}

func OauthUser(token string, service string) (int, error, string, time.Time) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", GetApiUrl(token, service), nil)
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "get api url with token "+string(token)+"and service "+string(service)), "", time.Time{}
	}

	if service == "yandex" {
		request.Header.Set("Authorization", "OAuth "+token)
	}
	response, err := client.Do(request)

	defer response.Body.Close()
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "cant HTTP request"), "", time.Time{}
	}

	contents, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "body ReadAll error"), "", time.Time{}
	}

	uuid, name, err := GetOauthUser(service, contents)
	if err != nil {
		return http.StatusForbidden, errors.Wrap(err, "err in get oauth user"), "", time.Time{}
	}
	fmt.Println(uuid)
	fmt.Println(name)
	searchingUser, err := user.GetUser(uuid)
	if err != nil {
		searchingUser, err = user.CreateUser(name, uuid, DefaultPassword)
		if err != nil {
			return http.StatusBadRequest, errors.Wrap(err, "err in user data"), "", time.Time{}
		}
	}

	randToken, expiration, err := session.SetToken(searchingUser.Id)
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "err set token"), "", time.Time{}
	}

	return 0, nil, randToken, expiration

}
