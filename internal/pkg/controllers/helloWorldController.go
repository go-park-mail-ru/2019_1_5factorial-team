package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/tempUsers"
	"net/http"
)

func HW(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("World"))
	tempUsers.AddUser(tempUsers.User{
		Email: "kek.k.ek",
		Nickname: "kek",
		Password: "password",
		Score: 100500,
		AvatarType: "jpg",
		AvatarLink: "./avatars/default.jpg"})

	tempUsers.PrintUsers()
}
