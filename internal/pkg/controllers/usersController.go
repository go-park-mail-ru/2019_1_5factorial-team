package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/tempUsers"
	"io/ioutil"
	"net/http"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"login":
// 	"password":
type SingUp struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func CreateUser(res http.ResponseWriter, req *http.Request)  {
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

	tempUsers.AddUser(tempUsers.User{

	})

}
