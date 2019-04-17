package random

import "math/rand"

func RandBool() bool {
	return rand.Intn(300) % 2 == 0
}
