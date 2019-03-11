package controllers

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"login":
// 	"password":
// пока только логин, без почты
type signInRequest struct {
	Login    string `json:"loginOrEmail"`
	Password string `json:"password"`
}

// SignIn godoc
// @Title Sign In
// @Summary Sign in with your account with email and password, set session cookie
// @ID post-session
// @Accept  json
// @Produce  json
// @Param AuthData body controllers.signInRequest true "user auth data"
// @Success 200 {object} OkResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /session [post]
func SignIn(res http.ResponseWriter, req *http.Request) {

	data := signInRequest{}
	status, err := ParseRequestIntoStruct(true, req, &data)
	if err != nil {
		ErrResponse(res, status, err.Error())

		log.Println(errors.Wrap(err, "ParseRequestIntoStruct error"))
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

	cookie := session.CreateHttpCookie(randToken, expiration)

	http.SetCookie(res, cookie)
	OkResponse(res, "ok auth")
}

// SignOut godoc
// @Title Sign Out
// @Summary Sign out from your account, expire cookie
// @ID delete-session
// @Produce  json
// @Success 200 {object} OkResponse
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Router /session [delete]
func SignOut(res http.ResponseWriter, req *http.Request) {

	currentSession, err := req.Cookie(session.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		return
	}

	err = session.DeleteToken(currentSession.Value)
	if err != nil && err.Error() == session.NoTokenFound {
		// bad token
		log.Println(errors.Wrap(err, "cannot delete token from current session, user cookie will set expired"))
	} else if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		log.Println(errors.Wrap(err, "delete token error"))
		return
	}

	// on other errors -> not logout, just answer ErrResponse

	status, err := DropUserCookie(res, req)
	if err != nil {
		ErrResponse(res, status, err.Error())

		log.Println(errors.Wrap(err, "cannot drop user cookie"))
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
		status, err := DropUserCookie(res, req)
		if err != nil {
			ErrResponse(res, status, err.Error())

			log.Println(errors.Wrap(err, "cannot drop user cookie"))
			return
		}

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
