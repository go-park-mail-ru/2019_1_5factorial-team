package oauth

import "strings"

func getLoginFromEmail(email string) string {
	return strings.Split(email, "@")[0]
}

func getEmailAndLoginFromVK(idVk string, nameVk string, lastNameVk string) (string, string) {
	email := idVk + "@vk.ru"
	login := nameVk + "_" + lastNameVk
	return email, login
}
