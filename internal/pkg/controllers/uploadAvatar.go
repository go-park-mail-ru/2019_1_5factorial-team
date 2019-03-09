package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/fileproc"
)

//const maxUploadSize = 2 * 1024 * 1024 // 2 mb
// const uploadPath = "../../src/avatars"

func UploadAvatar(res http.ResponseWriter, req *http.Request) {

	// проверка на максимально допустимый размер
	req.Body = http.MaxBytesReader(res, req.Body, fileproc.MaxUploadSize)
	if err := req.ParseMultipartForm(fileproc.MaxUploadSize); err != nil {
		ErrResponse(res, http.StatusBadRequest, "FILE_TOO_BIG")

		return
	}

	file, headers, err := req.FormFile("upload")
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "INVALID_FILE_IN_REQUEST")

		return
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "INVALID_FILE_CANT_READAll")

		return
	}

	//  я забыл какие у нас типы если что можно изи добавить

	filetype := http.DetectContentType(fileBytes)
	if !fileproc.CheckFileType(filetype) {
		ErrResponse(res, http.StatusBadRequest, "INVALID_FILE_TYPE")
		return
	}

	//  забиваю на имя фала генерю новое
	fileName := fileproc.RandToken(12)
	fileEndings := filepath.Ext(headers.Filename)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "CANT_READ_FILE_TYPE")

		return
	}
	newPath := filepath.Join(fileproc.UploadPath, fileName+fileEndings)
	fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

	// записываем файл
	newFile, err := os.Create(newPath)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "CANT_WRITE_FILE")

		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		ErrResponse(res, http.StatusInternalServerError, "CANT_WRITE_FILE")

		return
	}
	res.Write([]byte("SUCCESS"))

}
