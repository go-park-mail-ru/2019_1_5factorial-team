package session

import (
	"crypto/rand"
	"fmt"
)

func GenerateToken() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
