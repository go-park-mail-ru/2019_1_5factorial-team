package random

import (
	"github.com/google/uuid"
	"math/rand"
)

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	return id.String(), err
}