package gameLogic

import "math/rand"

type Symbol int

const (
	Top Symbol = iota
	Left
	Right
	Bottom
)

func GenerateSybolsSlice(len int) []Symbol {
	res := make([]Symbol, 0, len)
	for i := 0; i < len; i++ {
		res = append(res, Symbol(rand.Intn(4)))
	}

	return res
}
