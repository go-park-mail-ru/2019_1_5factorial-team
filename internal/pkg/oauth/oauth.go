package oauth

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
)

const DefaultPassword = ""
const PreCreateUserErrorString = "Invalid login"

type Token struct {
	Token string `json:"token"`
}

func OauthUser(token string, service string) (int, error, string, time.Time) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", GetApiUrl(token, service), nil)
	if err != nil {

		return http.StatusBadRequest, errors.Wrapf(err, "get api url with token %q and service %q", token, service), "", time.Time{}

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
	if response.StatusCode != http.StatusOK || err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "body ReadAll error"), "", time.Time{}
	}

	uuid, name, err := GetOauthUser(service, contents)
	if err != nil {
		return http.StatusForbidden, errors.Wrap(err, "err in get oauth user"), "", time.Time{}
	}
	// TODO(mrocumare) после прикручивания базы прописать GetUser и CreateUser
	searchingUser, err := user.IdentifyUser(uuid, DefaultPassword)
	if err != nil && errors.Cause(err).Error() == PreCreateUserErrorString {
		searchingUser, err = user.CreateUser(name, uuid, DefaultPassword)
		if err != nil {
			return http.StatusBadRequest, errors.Wrap(err, "err in user data"), "", time.Time{}
		}
	}

	randToken, expiration, err := session.SetToken(searchingUser.ID.Hex())
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "err set token"), "", time.Time{}
	}

	return 0, nil, randToken, expiration

}
