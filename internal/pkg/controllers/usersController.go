package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/grpcErr"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
		return http.StatusBadRequest, errors.New("already auth, ctx.authorized shouldn't be " + strconv.FormatBool(auth))
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
	currentSession, err := req.Cookie(config.Get().CookieConfig.CookieName)
	if err == http.ErrNoCookie {
		// бесполезная проверка, так кука валидна, но по гостайлу нужна

		return http.StatusUnauthorized, errors.Wrap(err, "not authorized")
	}

	currentSession.Expires = time.Unix(0, 0)
	http.SetCookie(res, currentSession)

	return 0, nil
}

// SignUp godoc
// @Title Sign Up
// @Summary Create account in our perfect game
// @ID sign-up
// @Produce json
// @Param AuthData body controllers.SingUpRequest true "user data to create"
// @Success 200 {string} ok response
// @Failure 400 {object} controllers.errorResponse
// @Failure 500 {object} controllers.errorResponse
// @Router /api/user [post]
func SignUp(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)
	ctx := context.Background()

	data := SingUpRequest{}
	statusErr, err := ParseRequestIntoStruct(true, req, &data)
	if err != nil {
		ErrResponse(res, statusErr, err.Error())

		ctxLogger.Error(errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}

	// TODO(smet1): валидация на данные, правда ли мыло - мыло, а самолет - вертолет?
	fmt.Println(data)

	u, err := authGRPC.CreateUser(ctx, &grpcAuth.UserNew{
		Nickname: data.Login,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		//ErrResponse(res, http.StatusConflict, err.Error())
		//
		//ctxLogger.Error(errors.Wrap(err, "err in user data"))
		//return
		st, ok := status.FromError(err)
		if !ok {
			ErrResponse(res, http.StatusInternalServerError, err.Error())

			ctxLogger.Error(errors.Wrap(err, "err in user data, cant convert err in status.FromError"))
			return
		}

		ErrResponse(res, grpcErr.GetHTTPStatus(st.Code()), st.Message())

		ctxLogger.Error(errors.Wrapf(err, "grpc code = %d, mes = %s", st.Code(), st.Message()))
		return

		//if errors.Cause(err).(*mgo.LastError).Code == MongoConflictCode {
		//	if strings.Contains(errors.Cause(err).(*mgo.LastError).Err, data.Login) {
		//		ErrResponse(res, http.StatusConflict, "login conflict")
		//
		//		ctxLogger.Error(errors.Wrap(err, "err in user data"))
		//		return
		//
		//	} else if strings.Contains(errors.Cause(err).(*mgo.LastError).Err, data.Email) {
		//		ErrResponse(res, http.StatusConflict, "email conflict")
		//
		//		ctxLogger.Error(errors.Wrap(err, "err in user data"))
		//		return
		//	}
		//}
	}

	cookieGRPC, err := authGRPC.CreateSession(ctx, &grpcAuth.UserID{ID: u.ID})
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		ctxLogger.Error(errors.Wrap(err, "Set token from grpc returned error"))
		return
	}

	timeCookie, err := time.Parse(time.RFC3339, cookieGRPC.Expiration)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		ctxLogger.Error(errors.Wrap(err, "cant convert time from string"))
		return
	}

	cookie := session.CreateHttpCookie(cookieGRPC.Token, timeCookie)

	http.SetCookie(res, cookie)
	OkResponse(res, "signUp ok")

	ctxLogger.Infof("OK response\n\t--id = %s,\n\t--nickname = %s,\n\t--email = %s,\n\t--score = %d",
		u.ID, u.Nickname, u.Email, u.Score)
	ctxLogger.Infof("OK set cookie\n\t--token = %s,\n\t--path = %s,\n\t--expires = %s,\n\t--httpOnly = %t",
		cookie.Value, cookie.Path, cookie.Expires, cookie.HttpOnly)
}

// GetUser godoc
// @Title get user
// @Summary Get user by id
// @ID get-user
// @Accept json
// @Produce json
// @Param id query int true "user id"
// @Success 200 {object} controllers.UserInfoResponse
// @Failure 400 {object} controllers.errorResponse
// @Failure 500 {object} controllers.errorResponse
// @Router /api/user/{id} [get]
func GetUser(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)
	ctx := context.Background()

	requestVariables := mux.Vars(req)
	if requestVariables == nil {
		ErrResponse(res, http.StatusBadRequest, "user id not provided")

		ctxLogger.Error(errors.New("no vars found"))
		return
	}

	searchingID, ok := requestVariables["id"]
	if !ok {
		ErrResponse(res, http.StatusInternalServerError, "error")

		ctxLogger.Error(errors.New("vars found, but cant found id"))
		return
	}

	searchingUser, err := authGRPC.GetUserByID(ctx, &grpcAuth.User{ID: searchingID})
	if err != nil {
		ErrResponse(res, http.StatusNotFound, "user with this id not found")

		ctxLogger.Error(errors.Wrap(err, "user with this id not found"))
		return
	}

	OkResponse(res, UserInfoResponse{
		Email:      searchingUser.Email,
		Nickname:   searchingUser.Nickname,
		Score:      int(searchingUser.Score),
		AvatarLink: searchingUser.AvatarLink,
	})

	ctxLogger.Infof("OK response\n\t--email = %v,\n\t--nickname = %v,\n\t--score = %v,\n\t--avatarLink = %v",
		searchingUser.Email, searchingUser.Nickname, searchingUser.Score, searchingUser.AvatarLink)
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

// UpdateProfile godoc
// @Title Update profile
// @Summary Update current profile (only avatar, only password or both)
// @ID update-profile
// @Accept json
// @Produce json
// @Param AuthData body controllers.ProfileUpdateRequest true "user data to update"
// @Success 200 {object} controllers.ProfileUpdateResponse
// @Failure 400 {object} controllers.errorResponse
// @Failure 500 {object} controllers.errorResponse
// @Router /api/user [put]
func UpdateProfile(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)
	ctx := context.Background()

	data := ProfileUpdateRequest{}
	status, err := ParseRequestIntoStruct(false, req, &data)
	if err != nil {
		ErrResponse(res, status, err.Error())

		ctxLogger.Error(errors.Wrap(err, "ParseRequestIntoStruct error"))
		return
	}

	userId := req.Context().Value("userID").(string)

	_, err = authGRPC.UpdateUser(ctx, &grpcAuth.UpdateUserReq{
		ID:          userId,
		NewAvatar:   data.Avatar,
		OldPassword: data.OldPassword,
		NewPassword: data.NewPassword,
	})

	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		ctxLogger.Error(errors.Wrap(err, "UpdateUser error"))
		return
	}

	u, err := authGRPC.GetUserByID(ctx, &grpcAuth.User{ID: userId})
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, err.Error())

		ctxLogger.Error(errors.Wrap(err, "UpdateUser error"))
		return
	}

	OkResponse(res, ProfileUpdateResponse{
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      int(u.Score),
		AvatarLink: u.AvatarLink,
	})

	ctxLogger.Infof("OK response\n\t--id = %s,\n\t--nickname = %s,\n\t--email = %s,\n\t--score = %d,\n\t--avatar = %s",
		u.ID, u.Nickname, u.Email, u.Score, u.AvatarLink)
}

type UsersCountInfoResponse struct {
	Count int `json:"count"`
}

// UsersCountInfo godoc
// @Title Get users count
// @Summary get count of registered users
// @ID get-users-count
// @Produce json
// @Success 200 {object} controllers.UsersCountInfoResponse
// @Router /api/user/count [get]
func UsersCountInfo(res http.ResponseWriter, req *http.Request) {
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
	authGRPC := req.Context().Value("authGRPC").(grpcAuth.AuthCheckerClient)
	ctx := context.Background()

	count, err := authGRPC.GetUsersCount(ctx, &grpcAuth.Nothing{})
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, err.Error())

		ctxLogger.Error(err.Error())
		return
	}

	OkResponse(res, UsersCountInfoResponse{
		Count: int(count.Count),
	})

	ctxLogger.Info("OK response, count = ", count)
}
