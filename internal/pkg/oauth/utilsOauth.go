package oauth

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type GoogleUser struct {
	Email string `json:"email"`
}

type VkUser struct {
	ID       int    `json:"id"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
}

type VkResponse struct {
	Users []VkUser `json:"response"`
}

type YandexUser struct {
	Email string `json:"default_email"`
}

func GetOauthUser(service string, contents []byte) (string, string, error) {

	switch service {
	case "vk":
		userInfo := VkResponse{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", errors.Wrap(err, "json parsing error")
		}

		if len(userInfo.Users) == 0 {
			return "", "", errors.New("not valid token")
		}
		vkInfo := (userInfo.Users)[0]
		vkEmail, vkLogin := getEmailAndLoginFromVK(strconv.Itoa(vkInfo.ID), vkInfo.Name, vkInfo.LastName)
		return vkEmail, vkLogin, nil
	case "yandex":
		userInfo := YandexUser{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", errors.Wrap(err, "json parsing error")
		}

		yandexNickname := getLoginFromEmail(userInfo.Email) + "_Y"
		return userInfo.Email, yandexNickname, nil
	case "google":
		userInfo := GoogleUser{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", errors.Wrap(err, "json parsing error")
		}
		googleNickname := getLoginFromEmail(userInfo.Email) + "_G"
		return userInfo.Email, googleNickname, nil
	default:
		err := errors.New("invalid service")
		return "", "", err
	}
}

func GetApiUrl(token string, service string) string {
	switch service {
	case "vk":
		return ("https://api.vk.com/method/users.get?fields=email,photo_50&access_token=" + token + "&v=5.52")
	case "google":
		return ("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	case "yandex":
		return "https://login.yandex.ru/info"
	default:
		return "error_service"
	}
}
