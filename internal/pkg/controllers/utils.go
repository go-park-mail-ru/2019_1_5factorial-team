package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type GoogleUser struct {
	Email string `json:"email"`
}

type VkUser struct {
	Response []struct {
		ID   int    `json:"id"`
		Name string `json:"first_name"`
	} `json:"response"`
}

type YandexUser struct {
	Email string `json:"default_email"`
}

//TODO(): разобраться с логированием и ошибками
func GetOauthUser(service string, contents []byte) (string, string, error) {
	if service == "vk" {
		userInfo := VkUser{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", err
		}
		return strconv.Itoa((userInfo.Response)[0].ID), (userInfo.Response)[0].Name, nil
	}
	if service == "yandex" {
		userInfo := YandexUser{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", err
		}
		return userInfo.Email, userInfo.Email, nil
	}
	if service == "google" {
		userInfo := GoogleUser{}
		err := json.Unmarshal(contents, &userInfo)
		if err != nil {
			return "", "", err
		}
		return userInfo.Email, userInfo.Email, nil
	}
	err := errors.New("invalid service")
	return "", "", err
	// err := json.Unmarshal(contents, &userInfo)
	// if err != nil {
	// 	return "", "", err
	// }
	// if service == "vk" {
	// return strconv.Itoa((userInfo.Response)[0].ID), (userInfo.Response)[0].Name, nil
	// }
	// return userInfo.Email, userInfo.Email, nil
}

func GetApiUrl(token string, service string) string {
	if service == "vk" {
		return ("https://api.vk.com/method/users.get?fields=email,photo_50&access_token=" + token + "&v=5.52")
	}
	if service == "yandex" {
		return "https://login.yandex.ru/info"
	}
	return ("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
}
