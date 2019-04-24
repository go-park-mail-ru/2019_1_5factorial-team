package controllers

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/validator"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
// @Success 200 {string} ok response
// @Failure 400 {object} controllers.errorResponse
// @Failure 500 {object} controllers.errorResponse
// @Router /session [post]
func SignIn(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	ctxLogger.Info("============================================")

	data := signInRequest{}
	status, err := ParseRequestIntoStruct(true, req, &data)
	if err != nil {
		ErrResponse(res, status, err.Error())

		ctxLogger.Error(errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}

	flagValidUser := validator.ValidLogin(data.Login, data.Password)
	if !flagValidUser {
		ErrResponse(res, http.StatusBadRequest, "invalid user data")
		ctxLogger.Error(errors.Wrap(err, "err in user data"))
		return
	}

	u, err := user.IdentifyUser(data.Login, data.Password)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "Wrong password or login")

		ctxLogger.Error(errors.Wrap(err, "Wrong password or login"))
		return
	}

	randToken, expiration, err := session.SetToken(u.ID.Hex())
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		ctxLogger.Error(errors.Wrap(err, "Set token returned error"))
		return
	}

	cookie := session.CreateHttpCookie(randToken, expiration)

	http.SetCookie(res, cookie)
	OkResponse(res, "ok auth")

	ctxLogger.Infof("OK response\n\t--id = %s,\n\t--nickname = %s,\n\t--email = %s,\n\t--score = %d",
		u.ID.Hex(), u.Nickname, u.Email, u.Score)
	ctxLogger.Infof("OK set cookie\n\t--token = %s,\n\t--path = %s,\n\t--expires = %s,\n\t--httpOnly = %t",
		cookie.Value, cookie.Path, cookie.Expires, cookie.HttpOnly)
}

// SignOut godoc
// @Title Sign Out
// @Summary Sign out from your account, expire cookie
// @ID delete-session
// @Produce  json
// @Success 200 {string} ok response
// @Failure 400 {object} controllers.errorResponse
// @Failure 401 {object} controllers.errorResponse
// @Router /session [delete]
func SignOut(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	ctxLogger.Info("===========================================")

	currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна
		ErrResponse(res, http.StatusUnauthorized, "not authorized")

		ctxLogger.Error(errors.Wrap(err, "not authorized"))
		return
	}

	err = session.DeleteToken(currentSession.Value)
	if err != nil && err.Error() == session.NoTokenFound {
		// bad token
		ctxLogger.Error(errors.Wrap(err, "cannot delete token from current session, user cookie will set expired"))
	} else if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		ctxLogger.Error(errors.Wrap(err, "delete token error"))
		return
	}

	// on other errors -> not logout, just answer ErrResponse

	status, err := DropUserCookie(res, req)
	if err != nil {
		ErrResponse(res, status, err.Error())

		ctxLogger.Error(errors.Wrap(err, "cannot drop user cookie"))
		return
	}

	OkResponse(res, "ok logout")

	ctxLogger.Info("OK response, cookie set expired, session deleted")
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

// GetUserFromSession godoc
// @Title get current user info
// @Summary Current user info
// @ID get-user-from-sesion
// @Produce json
// @Success 200 {object} controllers.UserInfoResponse
// @Failure 401 {object} controllers.errorResponse
// @Router /api/user [get]
func GetUserFromSession(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	ctxLogger.Info("===========================================")

	id := req.Context().Value("userID").(string)
	if id == "" {
		ErrResponse(res, http.StatusBadRequest, errors.New("invalid empty id").Error())

		ctxLogger.Error(errors.New("invalid empty id"))
		return
	}
	u, err := user.GetUserById(id)
	if err != nil {
		// проверка на невалидный айди юзера
		status, err := DropUserCookie(res, req)
		if err != nil {
			ErrResponse(res, status, err.Error())

			ctxLogger.Error(errors.Wrap(err, "cannot drop user cookie"))
			return
		}

		ErrResponse(res, http.StatusBadRequest, "error")

		ctxLogger.Error(errors.Wrap(err, "user have invalid id"))
		return
	}

	OkResponse(res, UserInfoResponse{
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      u.Score,
		AvatarLink: u.AvatarLink,
	})

	ctxLogger.Infof("OK response\n\t--email = %v,\n\t--nickname = %v,\n\t--score = %v,\n\t--avatarLink = %v",
		u.Email, u.Nickname, u.Score, u.AvatarLink)
}

// IsSessionValid godoc
// @Title check session
// @Summary check session of current user
// @ID is-session-valid
// @Produce json
// @Success 200 {string} ok message
// @Router /api/session [get]
func IsSessionValid(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	ctxLogger.Info("===========================================")

	OkResponse(res, "session is valid")
	ctxLogger.Info("session is valid")
}
