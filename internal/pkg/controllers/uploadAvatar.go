package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/sirupsen/logrus"

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
	ctxLogger := req.Context().Value("logger").(*logrus.Entry)

	// проверка на максимально допустимый размер
	req.Body = http.MaxBytesReader(res, req.Body, config.Get().StaticServerConfig.MaxUploadSize)
	if err := req.ParseMultipartForm(config.Get().StaticServerConfig.MaxUploadSize); err != nil {
		ErrResponse(res, http.StatusBadRequest, "file too big")

		ctxLogger.Error(errors.Wrap(err, "file too big"))
		return
	}

	file, headers, err := req.FormFile("upload")
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "invalid file in request")

		ctxLogger.Info(errors.Wrap(err, "invalid file in request"))
		return
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ErrResponse(res, http.StatusBadRequest, "invalid file cant ReadAll")

		ctxLogger.Info(errors.Wrap(err, "invalid file cant ReadAll"))
		return
	}

	filetype := http.DetectContentType(fileBytes)
	if !fileproc.CheckFileType(filetype) {
		ErrResponse(res, http.StatusBadRequest, "invalid file type")

		ctxLogger.Info(errors.Wrap(err, "invalid file type"))
		return
	}

	//  забиваю на имя файла генерю новое
	fileName := fileproc.RandToken(12)
	fileExtension := filepath.Ext(headers.Filename)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "cant read file type")

		ctxLogger.Info(errors.Wrap(err, "cant read file type"))
		return
	}

	// записываем файл
	resultFile, err := fileproc.CreateResultFile(fileName, fileExtension, filetype, fileBytes)
	if err != nil {
		ErrResponse(res, http.StatusInternalServerError, "cant write file")

		ctxLogger.Info(errors.Wrap(err, "cant write file"))
		return
	}

	OkResponse(res, AvatarLinkResponse{
		//AvatarLink: "/static/" + resultFile,
		AvatarLink: resultFile,
	})

	ctxLogger.Infof("OK response\n\t--avatar_link = /static/%s,\n\t--fileExtention = %s,\n\t--filetype = %s",
		resultFile, fileExtension, filetype)
}
