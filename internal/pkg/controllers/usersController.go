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
	"time"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"login":
//	"email":
// 	"password":
type SingUpRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id int64 `json:"id"`
}

func SignUp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("createUser")

	id := req.Context().Value("authorized").(bool)
	if id == true {
		ErrResponse(res, http.StatusBadRequest, "already auth")

		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "body parsing error")

		log.Println(errors.Wrap(err, "body parsing error"))
		return
	}

	fmt.Println()
	data := SingUpRequest{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "json parsing error")

		log.Println(errors.Wrap(err, "json parsing error"))
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
	expiration := time.Now().Add(10 * time.Hour)
	randToken := session.GenerateToken()

	err = session.SetToken(randToken, u.Id)
	// чет дичь
	if err != nil {
		for {
			randToken = session.GenerateToken()
			err = session.SetToken(randToken, u.Id)
			if err != nil {
				break
			}
		}
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   randToken,
		Expires: expiration,
		HttpOnly: true,
	}


	http.SetCookie(res, &cookie)
	OkResponse(res, "signUp ok")
}
