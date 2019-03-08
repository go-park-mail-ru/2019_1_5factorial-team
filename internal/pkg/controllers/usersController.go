package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func GetUser(res http.ResponseWriter, req *http.Request) {
	requestVariables := mux.Vars(req)
	if requestVariables == nil {
		ErrResponse(res, http.StatusBadRequest, "user id not provided")

		log.Println(errors.New("no vars found"))
		return
	}

	searchingID, ok := requestVariables["id"]
	if !ok {
		ErrResponse(res, http.StatusInternalServerError, "error")

		log.Println(errors.New("vars found, but cant found id"))
		return
	}

	intID, err := strconv.ParseInt(searchingID, 10, 64)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "bad id")

		log.Println(errors.New("cannot convert id from string"))
		return
	}

	searchingUser, err := user.GetUserById(intID)
	if err != nil {
		ErrResponse(res, http.StatusNotFound, "user with this id not found")

		log.Println(errors.Wrap(err, "404 error"))
		return
	}

	OkResponse(res, UserInfoResponse{
		Email:      searchingUser.Email,
		Nickname:   searchingUser.Nickname,
		Score:      searchingUser.Score,
		AvatarLink: searchingUser.AvatarLink,
	})
}

// 'Content-Type': 'application/json; charset=utf-8'
// 	"avatar":
//	"old_password":
// 	"new_password":
type ProfileUpdateRequest struct {
	Avatar      string `json:"avatar"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ProfileUpdateResponse struct {
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Score      int    `json:"score"`
	AvatarLink string `json:"avatar_link"`
}

func UpdateProfile(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "body parsing error")

		log.Println(errors.Wrap(err, "body parsing error"))
		return
	}

	fmt.Println()
	data := ProfileUpdateRequest{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "json parsing error")

		log.Println(errors.Wrap(err, "json parsing error"))
		return
	}

	err = user.UpdateUser(req.Context().Value("userID").(int64), data.Avatar, data.NewPassword, data.OldPassword)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println(errors.Wrap(err, "UpdateUser error"))
		return
	}

	u, _ := user.GetUserById(req.Context().Value("userID").(int64))

	OkResponse(res, ProfileUpdateResponse{
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      u.Score,
		AvatarLink: u.AvatarLink,
	})
}
