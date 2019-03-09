package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"login":
//	"email":
// 	"password":
type SingUpRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id int64 `json:"id"`
}

func ParseRequestIntoStruct(auth bool, req *http.Request, requestStruct interface{}) (int, error) {

	isAuth := req.Context().Value("authorized").(bool)
	if isAuth == auth {
		return http.StatusBadRequest, errors.New("already auth")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "body parsing error")
	}

	err = json.Unmarshal(body, &requestStruct)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "json parsing error")
	}

	return 0, nil
}

func SignUp(res http.ResponseWriter, req *http.Request) {

	data := SingUpRequest{}
	status, err := ParseRequestIntoStruct(true, req, &data)
	if err != nil {
		ErrResponse(res, status, err.Error())

		log.Println(errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}

	// TODO(smet1): валидация на данные, правда ли мыло - мыло, а самолет - вертолет?
	fmt.Println(data)

	u, err := user.CreateUser(data.Login, data.Email, data.Password)
	if err != nil {
		// TODO(smet1): указать точную ошибку
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println(errors.Wrap(err, "err in user data"))
		return
	}
	user.PrintUsers()

	// TODO(smet1): сразу его логинить или нет????
	randToken := session.GenerateToken()

	randToken, expiration, err := session.SetToken(u.Id)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    randToken,
		Expires:  expiration,
		HttpOnly: session.HttpOnly,
	}

	http.SetCookie(res, &cookie)
	OkResponse(res, "signUp ok")
}
