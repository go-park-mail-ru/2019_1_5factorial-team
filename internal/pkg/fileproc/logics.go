package fileproc

import (
	"crypto/rand"
	"fmt"
)

const MaxUploadSize = 2 * 1024 * 1024 // 2 mb
const UploadPath = "../../src/avatars"

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

func CheckFileType(a string) bool {
	arrayTypes := []string{"image/jpeg", "image/jpg", "image/png"}
	for _, b := range arrayTypes {
		if b == a {
			return true
		}
	}
	return false
}
