package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/fileproc"
	"github.com/pkg/errors"
)

// 'Content-Type': 'application/json; charset=utf-8'
// 	"avatarlink":
type AvatarLinkResponse struct {
	AvatarLink string `json:"avatar_link"`
}

// UploadAvatar godoc
// @Title Upload Avatar
// @Summary upload avatar on server
// @ID upload-avatar
// @Accept json
// @Produce json
// @Success 200 {object} controllers.AvatarLinkResponse
// @Failure 400 {object} controllers.errorResponse
// @Router /api/upload_avatar [post]
func UploadAvatar(res http.ResponseWriter, req *http.Request) {
	log.Println(req.URL, "UploadAvatar")

	// проверка на максимально допустимый размер
	req.Body = http.MaxBytesReader(res, req.Body, fileproc.MaxUploadSize)
	if err := req.ParseMultipartForm(fileproc.MaxUploadSize); err != nil {
		ErrResponse(res, http.StatusBadRequest, "file too big")

		log.Println(errors.Wrap(err, "file too big"))
		return
	}

	file, headers, err := req.FormFile("upload")
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "invalid file in request")

		log.Println(errors.Wrap(err, "invalid file in request"))
		return
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "invalid file cant ReadAll")

		log.Println(errors.Wrap(err, "invalid file cant ReadAll"))
		return
	}

	filetype := http.DetectContentType(fileBytes)
	if !fileproc.CheckFileType(filetype) {
		ErrResponse(res, http.StatusBadRequest, "invalid file type")

		log.Println(errors.Wrap(err, "invalid file type"))
		return
	}

	//  забиваю на имя фала генерю новое
	fileName := fileproc.RandToken(12)
	fileExtension := filepath.Ext(headers.Filename)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "cant read file type")

		log.Println(errors.Wrap(err, "cant read file type"))
		return
	}

	// записываем файл
	resultFile, err := fileproc.CreateResultFile(fileName, fileExtension, filetype, fileBytes)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "cant write file")

		log.Println(errors.Wrap(err, "cant write file"))
		return
	}

	log.Println("\t", req.URL, "OkResponse")
	log.Println("\t\tavatar_link:", resultFile)
	log.Println("\t\tfileExtension:", fileExtension, ", filetype:", filetype)

	OkResponse(res, AvatarLinkResponse{
		AvatarLink: resultFile,
	})
}
