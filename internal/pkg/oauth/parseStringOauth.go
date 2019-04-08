package oauth

import "strings"

func GetNickname(email string, service string) string {

	ending := ""
	switch service {
	case "yandex":

		ending = "_Y"
	case "google":
		ending = "_G"
	}
	nickname := ""
	splitEmail := strings.Split(email, "")
	for i := 0; i < len(splitEmail); i++ {
		if splitEmail[i] == "@" {
			return nickname + ending
		}
		nickname = nickname + splitEmail[i]
	}
	return nickname
}

func GetVkData(idVk string, nameVk string, lastNameVk string) (string, string) {
	email := idVk + "@vk.ru"
	login := nameVk + " " + lastNameVk
	return email, login
}
