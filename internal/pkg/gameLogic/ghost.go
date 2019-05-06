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
	// ширина спрайта на призраков и на игроков (одинаковые)
	DefaultSpriteWidth = 164

	AxisLen             = 1440
	PlayerLeftPosition  = (AxisLen - DefaultSpriteWidth) / 2
	PlayerRightPosition = AxisLen / 2

	DefaultRightPosition   = 1440
	DefaultLeftPosition    = 0
	DefaultMovementSpeed   = 80
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
		g.X = DefaultLeftPosition - DefaultSpriteWidth
		g.Speed = DefaultMovementSpeed
	} else {
		g.X = DefaultRightPosition + DefaultSpriteWidth
		g.Speed = -DefaultMovementSpeed
	}

	//if g.X > 0 {
	//	g.Speed = -DefaultMovementSpeed
	//} else {
	//	g.Speed = DefaultMovementSpeed
	//}

	return g
}

func NewRightGhost() Ghost {
	return Ghost{
		Damage:  DefaultDamage,
		Symbols: GenerateSymbolsSlice(DefaultLenSymbolsSlice),
		X:       DefaultRightPosition + DefaultSpriteWidth,
		Speed:   -DefaultMovementSpeed,
	}
}

func NewLeftGhost() Ghost {
	return Ghost{
		Damage:  DefaultDamage,
		Symbols: GenerateSymbolsSlice(DefaultLenSymbolsSlice),
		X:       DefaultLeftPosition - DefaultSpriteWidth,
		Speed:   DefaultMovementSpeed,
	}
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

func (gq *GhostQueue) PushBack(item Ghost) {
	gq.mu.Lock()
	gq.Items = append(gq.Items, item)
	gq.mu.Unlock()
}

func (gq *GhostQueue) PopBack() {
	gq.mu.Lock()
	gq.Items = gq.Items[:len(gq.Items)-1]
	gq.mu.Unlock()
}

func (gq *GhostQueue) AddNewGhost() {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	if len(gq.Items) == 0 {
		gq.Items = append(gq.Items, NewRandomGhost())
	} else {
		if gq.Items[0].Speed > 0 {
			gq.Items = append(gq.Items, NewRightGhost())
		} else {
			gq.Items = append(gq.Items, NewLeftGhost())
		}
	}
}

func (gq *GhostQueue) PopFront() Ghost {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	if len(gq.Items) == 0 {
		return Ghost{}
	}

	retVal := gq.Items[0]
	gq.Items = gq.Items[1:]
	return retVal
}

// return true if first ghost reach player
func (gq *GhostQueue) MoveAllGhosts() (hit bool) {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	if len(gq.Items) == 0 {
		return
	}

	for i := 0; i < len(gq.Items); i++ {
		gq.Items[i].Move()
	}

	// коллизии, как хотела Надя
	if gq.Items[0].Speed > 0 && gq.Items[0].X >= PlayerLeftPosition-DefaultSpriteWidth {
		hit = true
	} else if gq.Items[0].Speed < 0 && gq.Items[0].X <= PlayerRightPosition+DefaultSpriteWidth {
		hit = true
	}

	//return gq.Items[0].X == (DefaultRightPosition-DefaultLeftPosition)/2
	return
}

func (gq *GhostQueue) PopSymbol(sym Symbol) int {
	gq.mu.Lock()
	defer gq.mu.Unlock()
	score := 0

	if len(gq.Items) == 0 {
		log.Println("len = 0, ret score = 0")

		return 0
	}

	newItems := make([]Ghost, 0, 1)
	for i := range gq.Items {
		log.Println(gq.Items[i])
		if len(gq.Items[i].Symbols) == 0 {
			log.Println("bug")
			continue
		}

		if gq.Items[i].Symbols[0] == sym {
			gq.Items[i].Symbols = gq.Items[i].Symbols[1:]

			score += ScoreMatchSymbol
		}

		if len(gq.Items[i].Symbols) != 0 {
			newItems = append(newItems, gq.Items[i])
		} else {
			score += ScoreKillGhost
		}
	}

	if len(newItems) != 0 {
		gq.Items = newItems
	}

	return score
}

func (gq *GhostQueue) Len() int {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	return len(gq.Items)
}
