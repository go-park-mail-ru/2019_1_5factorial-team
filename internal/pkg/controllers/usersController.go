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
	"time"
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

func DropUserCookie(res http.ResponseWriter, req *http.Request) (int, error) {
	currentSession, err := req.Cookie(session.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна

		return http.StatusUnauthorized, errors.Wrap(err, "not authorized")
	}

	currentSession.Expires = time.Unix(0, 0)
	http.SetCookie(res, currentSession)

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
		ErrResponse(res, http.StatusBadRequest, err.Error())

		log.Println(errors.Wrap(err, "err in user data"))
		return
	}
	user.PrintUsers()

	randToken, expiration, err := session.SetToken(u.Id)

	cookie := session.CreateHttpCookie(randToken, expiration)

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

// 'Content-Type': 'application/json; charset=utf-8'
// 	"email":
//	"nickname":
// 	"score":
//	"avatar_link":
type ProfileUpdateResponse struct {
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Score      int    `json:"score"`
	AvatarLink string `json:"avatar_link"`
}

func UpdateProfile(res http.ResponseWriter, req *http.Request) {
	// TODO(): использовать метод из дева для заполнения структуры из запроса
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "body parsing error")

		log.Println(errors.Wrap(err, "body parsing error"))
		return
	}

	// TODO(): аватарки должны будут обновлять по-другому, жду лелю
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

type UsersCountInfoResponse struct {
	Count int `json:"count"`
}

func UsersCountInfo(res http.ResponseWriter, req *http.Request) {
	count := user.GetUsersCount()
	OkResponse(res, UsersCountInfoResponse{
		Count: count,
	})
}
