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
	Login    string `json:"loginOrEmail"`
	Password string `json:"password"`
}
// пока только логин, без почты

func SignIn(res http.ResponseWriter, req *http.Request) {
	isAuth := req.Context().Value("authorized").(bool)
	if isAuth == true {
		ErrResponse(res, http.StatusBadRequest, "already auth")

		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "body parsing error")

		log.Println(errors.Wrap(err, "body parsing error"))
		return
	}
	defer req.Body.Close()

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

	randToken, expiration, err := session.SetToken(u.Id)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    randToken,
		Expires:  expiration,
		HttpOnly: session.HttpOnly,
	}

	http.SetCookie(res, &cookie)
	OkResponse(res, "ok auth")
}

func SignOut(res http.ResponseWriter, req *http.Request) {

	currentSession, err := req.Cookie("token")
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		return
	}

	currentSession.Expires = time.Unix(0, 0)
	http.SetCookie(res, currentSession)

	err = session.DeleteToken(currentSession.Value)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "error")
		log.Println(errors.Wrap(err, "cannot delete token from current session, user cookie set expired"))

		return
	}

	OkResponse(res, "ok logout")
}

// 'Content-Type': 'application/json; charset=utf-8'
// 	"email":
// 	"nickname":
// 	"score":
// 	"avatar_link":

type UserInfoResponse struct {
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Score      int    `json:"score"`
	AvatarLink string `json:"avatar_link"`
}

func GetUserFromSession(res http.ResponseWriter, req *http.Request) {

	id := req.Context().Value("userID").(int64)
	u, err := user.GetUserById(id)
	if err != nil {
		// проверка на невалидный айди юзера
		currentSession, err := req.Cookie("token")
		currentSession.Expires = time.Unix(0, 0)
		http.SetCookie(res, currentSession)

		ErrResponse(res, http.StatusBadRequest, "error")
		log.Println(errors.Wrap(err, "user have invalid id"))

		return
	}

	OkResponse(res, UserInfoResponse{
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      u.Score,
		AvatarLink: u.AvatarLink,
	})
}

func IsSessionValid(res http.ResponseWriter, req *http.Request) {

	OkResponse(res, "session is valid")
}
