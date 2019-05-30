package controllers

import (
	"net/http"
	"strings"
	"testing"
)

var testsUsersCountInfo = []TestCases{
	{
		routerPath:     "/api/user/count",
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"count":21}`,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user/count",
		method:         "GET",
		url:            "/api/user/count",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"count":21}`,
		expectedStatus: http.StatusOK,
		authCtx:        true,
	},
}

func TestUsersCountInfo(t *testing.T) {
	//MainInit()
	err := testHandler(UsersCountInfo, testsUsersCountInfo, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsSignUp = []TestCases{
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: parse error: syntax error near offset 26 of '{\"email\": \"kek@email.kek\",}'"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"email": "kek@email.kek",}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth, ctx.authorized shouldn't be true"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"invalid avatar"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek", "email": "kek@email.kek","password": "password", "avatar_link":"kek"}`),
		urlValues:      "",
		expectedRes:    `{"error":"invalid avatar"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek1", "email": "kek@email1.kek","password": "password", "avatar_link":"001-default-avatar"}`),
		urlValues:      "",
		expectedRes:    ``,
		expectedStatus: http.StatusOK,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "POST",
		url:            "/api/user",
		body:           strings.NewReader(`{"login": "kekkekkek123", "email": "kek@email1.kek","password": "password", "avatar_link":"001-default-avatar"}`),
		urlValues:      "",
		expectedRes:    `{"error":"email conflict"}`,
		expectedStatus: http.StatusConflict,
		authCtx:        false,
	},
}

func TestSignUp(t *testing.T) {
	err := testHandler(SignUp, testsSignUp, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsGetUser = []TestCases{
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/0",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user with this id not found"}`,
		expectedStatus: http.StatusNotFound,
	},
	{
		routerPath:     "/api/user/{id:[0-9]+}",
		method:         "GET",
		url:            "/api/user/500000000000000000000000000000000000000000000000",
		body:           nil,
		urlValues:      "",
		expectedRes:    `{"error":"user with this id not found"}`,
		expectedStatus: http.StatusNotFound,
	},
}

func TestGetUser(t *testing.T) {
	//MainInit()
	err := testHandler(GetUser, testsGetUser, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var testsUpdateProfile = []TestCases{
	{
		routerPath:     "/api/user",
		method:         "PUT",
		url:            "/api/user",
		body:           strings.NewReader(`{"avatar": "kekkekkek1", "old_password": "1","new_password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"already auth, ctx.authorized shouldn't be false"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        false,
	},
	{
		routerPath:     "/api/user",
		method:         "PUT",
		url:            "/api/user",
		body:           strings.NewReader(`{"avatar": "kekkekkek1", "old_password": "1","new_password": "password"`),
		urlValues:      "",
		expectedRes:    `{"error":"json parsing error: EOF"}`,
		expectedStatus: http.StatusInternalServerError,
		authCtx:        true,
	},
	{
		routerPath:     "/api/user",
		method:         "PUT",
		url:            "/api/user",
		body:           strings.NewReader(`{"avatar": "kekkekkek1", "old_password": "1","new_password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"rpc error: code = Unknown desc = update user error: id isn't mongo's hex"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
	},
	{
		routerPath:     "/api/user",
		method:         "PUT",
		url:            "/api/user",
		body:           strings.NewReader(`{"avatar": "kekkekkek1", "old_password": "1","new_password": "password"}`),
		urlValues:      "",
		expectedRes:    `{"error":"rpc error: code = Unknown desc = update user error: user with this id not found"}`,
		expectedStatus: http.StatusBadRequest,
		authCtx:        true,
		userIDCtx:      "5556c0d9b49cd4582aaad41c",
	},
}

func TestUpdateProfile(t *testing.T) {
	err := testHandler(UpdateProfile, testsUpdateProfile, t)
	if err != nil {
		t.Errorf(err.Error())
	}
}
