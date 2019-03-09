package fileproc

import (
	"crypto/rand"
	"fmt"
)

//const maxUploadSize = 2 * 1024 * 1024 // 2 mb
//const uploadPath = "../../src/avatars"

func uploadPath() string {
	return "../../src/avatars"
}

func maxUploadSize() int {
	return 2 * 1024 * 1024 // 2 mb
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func checkFileType(a string) bool {
	arrayTypes := []string{"image/jpeg", "image/jpg", "image/png"}
	for _, b := range arrayTypes {
		if b == a {
			return true
		}
	}
	return false
}
