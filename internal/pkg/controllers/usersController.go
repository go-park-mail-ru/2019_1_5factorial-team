package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"io/ioutil"
	"net/http"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"login":
//	"email":
// 	"password":
type SingUp struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpResp struct {
	Id int `json:"id"`
}

func SignUp(res http.ResponseWriter, req *http.Request)  {
	fmt.Println("createUser")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// bad request
		fmt.Println(err)
		return
	}
	fmt.Println()
	data := SingUp{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		// json unmarshal error
		fmt.Println(err)
		return
	}
	fmt.Println(data)

	u, err := user.CreateUser(data.Login, data.Email, data.Password)
	if err != nil {
		// some errors with validation
		fmt.Println(err)
		return
	}
	user.PrintUsers()

	// return user id
	AddOkHeader(res)
	AddBody(res, SignUpResp{u.Id})
}
