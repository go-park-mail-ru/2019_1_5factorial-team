package controllers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	fileUrl := "https://golangcode.com/images/avatar.jpg"
	test_file_name := name + ".jpg"

	if err := DownloadFile(test_file_name, fileUrl); err != nil {
		panic(err)
	}

	tmpl, _ := template.ParseFiles("/Users/Mr.ocumare/Desktop/TP_2sem/GoLang/Project/2019_1_5factorial-team/src/templates/get_img.html")
	tmpl.Execute(res, data)
}

func DownloadFile(fileName string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join("../../src/img", filepath.Base(fileName)))
	//out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
