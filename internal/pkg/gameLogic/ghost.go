package gameLogic

import (
	"fmt"
	"sync"
)

type Ghost struct {
	X       int      `json:"x"`
	Speed   int      `json:"speed"`
	Damage  uint32   `json:"damage"`
	Sprite  string   `json:"sprite"`
	Symbols []Symbol `json:"symbols"`
}

func NewGhost(startPosition int, damage uint32, sprite string, speed int, symbolsLen int) Ghost {
	g := Ghost{
		X:       startPosition,
		Damage:  damage,
		Sprite:  sprite,
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

// стэк для удобства работы с призраками
type GhostQueue struct {
	items []Ghost
	mu    *sync.Mutex
}

func NewGhostStack() *GhostQueue {
	return &GhostQueue{
		items: make([]Ghost, 0, 1),
		mu:    &sync.Mutex{},
	}
}

func (gs *GhostQueue) PushBack(item Ghost) {
	gs.mu.Lock()
	gs.items = append(gs.items, item)
	gs.mu.Unlock()
}

func (gs *GhostQueue) PopBack() {
	gs.mu.Lock()
	gs.items = gs.items[:len(gs.items)-1]
	gs.mu.Unlock()
}

func (gs *GhostQueue) PopFront() {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if len(gs.items) == 0 {
		return
	}
	gs.items = gs.items[1:]
}

// return true if first ghost reach player
func (gs *GhostQueue) MoveAllGhosts() bool {
	gs.mu.Lock()
	fmt.Println("MoveAllGhosts lock")
	defer gs.mu.Unlock()

	if len(gs.items) == 0 {
		return false
	}

	for i := 0; i < len(gs.items); i++ {
		gs.items[i].Move()
		fmt.Printf("%d moved\n", i)
	}

	fmt.Println("MoveAllGhosts unlock")
	return gs.items[0].X == 0
}

func (gs *GhostQueue) Len() int {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	return len(gs.items)
}
