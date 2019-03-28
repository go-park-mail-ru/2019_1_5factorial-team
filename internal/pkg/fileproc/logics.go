package fileproc

import (
	"crypto/rand"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/server"
	"os"
	"path/filepath"
)

func RandToken(len int) string {
	token := make([]byte, len)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}

func CheckFileType(receivedFileType string) bool {
	arrayTypes := []string{"image/jpeg", "image/jpg", "image/png"}
	for _, arrayTypesCell := range arrayTypes {
		if arrayTypesCell == receivedFileType {
			return true
		}
	}
	return false
}

func CreateNewFile(fileName string, fileExtension string, filetype string) string {

	newFile := filepath.Join(
		server.GetInstance().StaticServerConfig.UploadPath,
		fileName+fileExtension,
		)
	//log.Printf("filetype: %s, file: %s\n", filetype, newFile)
	return newFile
}

func CreateResultFile(fileName string, fileExtension string, filetype string, fileBytes []byte) (string, error) {
	newFile, err := os.Create(CreateNewFile(fileName, fileExtension, filetype))
	if err != nil {
		return "", err
	}

	defer newFile.Close()
	_, err = newFile.Write(fileBytes)
	if err != nil || newFile.Close() != nil {
		return "", err
	}
	return fileName + fileExtension, nil
}
