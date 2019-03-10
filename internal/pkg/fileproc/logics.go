package fileproc

import (
	"crypto/rand"
	"fmt"
)

// func UploadPath() string {
// 	return "../../src/avatars"
// }

// func MaxUploadSize() int {
// 	return 2 * 1024 * 1024 // 2 mb
// }

func RandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func CheckFileType(getTypeFile string) bool {
	arrayTypes := []string{"image/jpeg", "image/jpg", "image/png"}
	for _, trueTypeFile := range arrayTypes {
		if trueTypeFile == getTypeFile {
			return true
		}
	}
	return false
}
