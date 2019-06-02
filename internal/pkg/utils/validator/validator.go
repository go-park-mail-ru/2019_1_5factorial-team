package validator

import "regexp"

const ExpEmail = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
const ExpLoginPassword = `^[a-z0-9._-]{4,20}$`
const ExpAvatar = `^\/[0-9]{3}-default-avatar.png$`

var (
	expEmail       = regexp.MustCompile(ExpEmail)
	expLopPassword = regexp.MustCompile(ExpLoginPassword)
	expAvatar      = regexp.MustCompile(ExpAvatar)
)

func ValidNewUser(login string, email string, password string) bool {
	//expEmail := regexp.MustCompile(ExpEmail)
	//expLopPassw := regexp.MustCompile(ExpLoginPassword)
	if expEmail.MatchString(email) && expLopPassword.MatchString(login) && expLopPassword.MatchString(password) {
		return true
	}
	return false
}

func ValidateAvatarDefault(avatar string) bool {
	//expAvatar := regexp.MustCompile(ExpAvatar)
	return expAvatar.MatchString(avatar)
}

func ValidUpdatePassword(password string) bool {
	//expLopPassw := regexp.MustCompile(ExpLoginPassword)

	return expLopPassword.MatchString(password)
}

func ValidLogin(login string, password string) bool {
	//expLopPassw := regexp.MustCompile(ExpLoginPassword)
	if expLopPassword.MatchString(login) && expLopPassword.MatchString(password) {
		return true
	}
	return false
}
