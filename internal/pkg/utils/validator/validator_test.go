package validator

import (
	"testing"
)

var validatePassword = []struct {
	password string
	want     bool
}{
	{"", false},
}

func TestValidUpdatePassword(t *testing.T) {
	for _, val := range validatePassword {
		res := ValidUpdatePassword(val.password)
		if res != val.want {
			t.Error("expected", val.want, "have", res)
		}
	}
}

var validateUser = []struct {
	login    string
	email    string
	password string
	want     bool
}{
	{"", "", "", false},
	{"kekkekkek", "kek@ke.k", "///////", false},
}

func TestValidNewUser(t *testing.T) {
	for _, val := range validateUser {
		res := ValidNewUser(val.login, val.email, val.password)
		if res != val.want {
			t.Error("expected", val.want, "have", res)
		}
	}
}

var validateLogin = []struct {
	login    string
	password string
	want     bool
}{
	{"", "", false},
	{"kekkekkek", "///////", false},
	{"kekkekkek", "password", true},
}

func TestValidLogin(t *testing.T) {
	for _, val := range validateLogin {
		res := ValidLogin(val.login, val.password)
		if res != val.want {
			t.Error("expected", val.want, "have", res)
		}
	}
}
