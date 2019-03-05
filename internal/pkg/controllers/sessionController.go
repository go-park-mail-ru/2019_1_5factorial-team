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
// 	"password":
type signInRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func SignIn(res http.ResponseWriter, req *http.Request) {
	//_, err := req.Cookie("user_id")
	//if err != http.ErrNoCookie {
	//	ErrResponse(res, http.StatusBadRequest, "already auth")
	//
	//	return
	//}
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

	data := signInRequest{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "json parsing error")

		log.Println(errors.Wrap(err, "json parsing error"))
		return
	}

	u, err := user.IdentifyUser(data.Login, data.Password)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "Wrong password or login")
		log.Println(errors.Wrap(err, "Wrong password or login"))
		return
	}

	fmt.Println(u)
	// unsafe
	// можно генерить вместо айди уникальный токен и пихать его в отдельную таблицу, которую постоянно проверять
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
	OkResponse(res, "")
}

func SignOut(res http.ResponseWriter, req *http.Request) {
	id := req.Context().Value("authorized").(bool)
	if id == false {
		ErrResponse(res, http.StatusBadRequest, "not authorized")

		return
	}

	currentSession, err := req.Cookie("token")
	if err == http.ErrNoCookie {
		ErrResponse(res, http.StatusBadRequest, "not authorized")

		return
	}

	currentSession.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(res, currentSession)

	err = session.DeleteToken(currentSession.Value)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "error")
		log.Println(errors.Wrap(err, "cannot delete token from current session, user cookie set expired"))

		return
	}

	OkResponse(res, "ok logout")
}

type userInfoResponse struct {
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Score int `json:"score"`
	AvatarLink string `json:"avatar_link"`
}

func GetUserFromSession(res http.ResponseWriter, req *http.Request) {

}