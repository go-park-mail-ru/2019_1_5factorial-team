package gameLogic

import "math/rand"

type Symbol int

const (
	BottomTriangle Symbol = iota
	TopTriangle
	Square
	Circle
	VertLine
	HorizLine
)

func GenerateSybolsSlice(len int) []Symbol {
	res := make([]Symbol, 0, len)
	for i := 0; i < len; i++ {
		res = append(res, Symbol(rand.Intn(6)))
	}

	return res
}
