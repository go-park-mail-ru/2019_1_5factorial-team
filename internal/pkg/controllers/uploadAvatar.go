package controllers

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "../../src/avatars"

func UploadAvatar(res http.ResponseWriter, req *http.Request) {

	// проверка на максимально допустимый размер
	req.Body = http.MaxBytesReader(res, req.Body, maxUploadSize)
	if err := req.ParseMultipartForm(maxUploadSize); err != nil {
		renderError(res, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	//fileType := req.PostFormValue("type")
	file, _, err := req.FormFile("upload")
	if err != nil {
		renderError(res, "INVALID_FILE_BIBA", http.StatusBadRequest)
		return
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(res, "INVALID_FILE_BOBA", http.StatusBadRequest)
		return
	}

	//  я забыл какие у нас типы если что можно изи добавить
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/png":
		break
	default:
		renderError(res, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	//  забиваю на имя фала генерю новое
	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(filetype)
	if err != nil {
		renderError(res, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return
	}
	newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
	fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

	// записываем файл
	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(res, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(res, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	res.Write([]byte("SUCCESS"))

}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
