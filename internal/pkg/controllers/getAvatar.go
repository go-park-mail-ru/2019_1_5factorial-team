package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetImg(writer http.ResponseWriter, request *http.Request) {

	// Filename := request.URL.Query().Get("file")  чет я пытался но у меня лапки

	Filename := request.FormValue("file")
	if Filename == "" {
		http.Error(writer, "GET_'file'_NOT_SPECIFIED_IN_URL.", 400)
		return
	}
	fmt.Println("CLIENT_REQUEST: " + Filename)

	Filename = "../../src/avatars/" + Filename

	Openfile, err := os.Open(Filename)
	defer Openfile.Close()
	if err != nil {
		http.Error(writer, "FILE_NOT_FOUND", 404)
		return
	}

	FileHeader := make([]byte, 2*1024*1024)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)

	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile)
	return
}
