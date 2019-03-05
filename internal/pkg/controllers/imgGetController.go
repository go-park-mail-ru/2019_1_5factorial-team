package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	img_name string
}

func GetImg(res http.ResponseWriter, req *http.Request) {
	data := ViewData{
		img_name: "Users List",
	}

	err := req.ParseForm()
	if err != nil {

	}

	name := req.FormValue("name")
	fmt.Print(name)

	// fileUrl := ""

	tmpl, _ := template.ParseFiles("/Users/Mr.ocumare/Desktop/TP_2sem/GoLang/Project/2019_1_5factorial-team/src/templates/get_img.html")
	tmpl.Execute(res, data)
}
