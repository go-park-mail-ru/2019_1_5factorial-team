package gameLogic

import (
	"errors"
	"math/rand"
)

type Symbol int

const (
	Up Symbol = iota
	Left
	Right
	Down
)

var symbolsDict = map[string]Symbol{
	"up":    Up,
	"left":  Left,
	"right": Right,
	"down":  Down,
}

func GenerateSymbolsSlice(len int) []Symbol {
	res := make([]Symbol, 0, len)
	for i := 0; i < len; i++ {
		res = append(res, Symbol(rand.Intn(4)))
	}

	return res
}

func MatchSymbol(in string) (Symbol, error) {
	if el, ok := symbolsDict[in]; ok {
		return el, nil
	}

	return -1, errors.New("not correct input")
}
