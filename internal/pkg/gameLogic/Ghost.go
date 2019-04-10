package gameLogic

type Ghost struct {
	X       int   `json:"x"`
	Speed   int   `json:"speed"`
	Damage  uint32   `json:"damage"`
	Sprite  string   `json:"sprite"`
	Symbols []Symbol `json:"symbols"`
}

func NewGhost(startPosition int, damage uint32, sprite string, speed int, symbolsLen int) Ghost {
	g := Ghost{
		X:      startPosition,
		Damage: damage,
		Sprite: sprite,
		Symbols: GenerateSybolsSlice(symbolsLen),
	}

	if g.X > 0 {
		g.Speed = -speed
	} else {
		g.Speed = speed
	}

	return g
}

func (gh *Ghost) Move() {
	gh.X += gh.Speed
}
