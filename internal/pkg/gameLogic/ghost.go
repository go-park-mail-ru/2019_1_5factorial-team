package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/random"
	"log"
	"sync"
)

type Ghost struct {
	X       int      `json:"x"`
	Speed   int      `json:"speed"`
	Damage  uint32   `json:"damage"`
	Symbols []Symbol `json:"symbols"`
}

const (
	DefaultRightPosition   = 1440
	DefaultLeftPosition    = 0
	DefaultMovementSpeed   = 100
	DefaultLenSymbolsSlice = 4
	DefaultDamage          = 1

	// за 1 призрака при 4 символах, можно получить 100
	ScoreKillGhost   = 60
	ScoreMatchSymbol = 10
)

func NewGhost(startPosition int, damage uint32, speed int, symbolsLen int) Ghost {
	g := Ghost{
		X:       startPosition,
		Damage:  damage,
		Symbols: GenerateSymbolsSlice(symbolsLen),
	}

	if g.X > 0 {
		g.Speed = -speed
	} else {
		g.Speed = speed
	}

	return g
}

func NewRandomGhost() Ghost {
	g := Ghost{
		//X:       DefaultRightPosition,
		Damage:  DefaultDamage,
		Symbols: GenerateSymbolsSlice(DefaultLenSymbolsSlice),
	}

	if random.RandBool() {
		g.X = DefaultLeftPosition
		g.Speed = DefaultMovementSpeed
	} else {
		g.X = DefaultRightPosition
		g.Speed = -DefaultMovementSpeed

	}

	//if g.X > 0 {
	//	g.Speed = -DefaultMovementSpeed
	//} else {
	//	g.Speed = DefaultMovementSpeed
	//}

	return g
}

func (gh *Ghost) Move() {
	gh.X += gh.Speed
}

// стэк для удобства работы с призраками
type GhostQueue struct {
	Items []Ghost `json:"items"`
	mu    *sync.Mutex
}

func NewGhostStack() *GhostQueue {
	return &GhostQueue{
		Items: make([]Ghost, 0, 1),
		mu:    &sync.Mutex{},
	}
}

func (gs *GhostQueue) PushBack(item Ghost) {
	gs.mu.Lock()
	gs.Items = append(gs.Items, item)
	gs.mu.Unlock()
}

func (gs *GhostQueue) PopBack() {
	gs.mu.Lock()
	gs.Items = gs.Items[:len(gs.Items)-1]
	gs.mu.Unlock()
}

func (gs *GhostQueue) PopFront() {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if len(gs.Items) == 0 {
		return
	}
	gs.Items = gs.Items[1:]
}

// return true if first ghost reach player
func (gs *GhostQueue) MoveAllGhosts() bool {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if len(gs.Items) == 0 {
		return false
	}

	for i := 0; i < len(gs.Items); i++ {
		gs.Items[i].Move()
	}

	return gs.Items[0].X == (DefaultRightPosition-DefaultLeftPosition)/2
}

func (gs *GhostQueue) PopSymbol(sym Symbol) int {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	score := 0

	if len(gs.Items) == 0 {
		log.Println("len = 0, ret score = 0")

		return 0
	}

	newItems := make([]Ghost, 0, 1)
	for i := range gs.Items {
		if gs.Items[i].Symbols[0] == sym {
			gs.Items[i].Symbols = gs.Items[i].Symbols[1:]

			score += ScoreMatchSymbol
		}
		if len(gs.Items[i].Symbols) != 0 {
			newItems = append(newItems, gs.Items[i])
		} else {
			score += ScoreKillGhost
		}
	}

	if len(newItems) != 0 {
		gs.Items = newItems
	}

	return score
}

func (gs *GhostQueue) Len() int {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	return len(gs.Items)
}
