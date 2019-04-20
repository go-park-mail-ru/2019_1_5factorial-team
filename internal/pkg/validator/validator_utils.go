package validator

import "regexp"

const _EXP_EMAIL = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
const _EXP_LOGIN_PASSWORD = `^[a-z0-9._-]{4,20}$`

func ValidNewUser(login string, email string, password string) bool {

	expEmail := regexp.MustCompile(_EXP_EMAIL)
	expLopPassw := regexp.MustCompile(_EXP_LOGIN_PASSWORD)
	if expEmail.MatchString(email) && expLopPassw.MatchString(login) && expLopPassw.MatchString(password) {
		return true
	}
	return false
}
func ValidUpdatePassword(password string) bool {
	expLopPassw := regexp.MustCompile(_EXP_LOGIN_PASSWORD)
	return expLopPassw.MatchString(password)
}
