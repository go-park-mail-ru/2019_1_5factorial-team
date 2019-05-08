package validator

import "regexp"

const ExpEmail = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
const ExpLoginPassword = `^[a-z0-9._-]{4,20}$`

func ValidNewUser(login string, email string, password string) bool {
	expEmail := regexp.MustCompile(ExpEmail)
	expLopPassw := regexp.MustCompile(ExpLoginPassword)
	if expEmail.MatchString(email) && expLopPassw.MatchString(login) && expLopPassw.MatchString(password) {
		return true
	}
	return false
}

func ValidUpdatePassword(password string) bool {
	expLopPassw := regexp.MustCompile(ExpLoginPassword)

	return expLopPassw.MatchString(password)
}
func ValidLogin(login string, password string) bool {
	expLopPassw := regexp.MustCompile(ExpLoginPassword)
	if expLopPassw.MatchString(login) && expLopPassw.MatchString(password) {
		return true
	}
	return false
}
