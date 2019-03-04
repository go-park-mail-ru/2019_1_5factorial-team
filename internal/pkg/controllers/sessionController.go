package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	_, err := req.Cookie("user_id")
	if err != http.ErrNoCookie {
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
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "user_id",
		Value:   strconv.Itoa(u.Id),
		Expires: expiration,
		HttpOnly: true,
	}

	http.SetCookie(res, &cookie)
	OkResponse(res, "")
}
