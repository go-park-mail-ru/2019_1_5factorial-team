package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
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
		Damage:  config.Get().GameConfig.DefaultDamage,
		Symbols: GenerateSymbolsSlice(config.Get().GameConfig.DefaultLenSymbolsSlice),
	}

	if random.RandBool() {
		g.X = config.Get().GameConfig.DefaultLeftPosition - config.Get().GameConfig.DefaultSpriteWidth
		g.Speed = config.Get().GameConfig.DefaultMovementSpeed
	} else {
		g.X = config.Get().GameConfig.DefaultRightPosition + config.Get().GameConfig.DefaultSpriteWidth
		g.Speed = -config.Get().GameConfig.DefaultMovementSpeed
	}

	return g
}

func NewRightGhost() Ghost {
	return Ghost{
		Damage:  config.Get().GameConfig.DefaultDamage,
		Symbols: GenerateSymbolsSlice(config.Get().GameConfig.DefaultLenSymbolsSlice),
		X:       config.Get().GameConfig.DefaultRightPosition + config.Get().GameConfig.DefaultSpriteWidth,
		Speed:   -config.Get().GameConfig.DefaultMovementSpeed,
	}
}

func NewLeftGhost() Ghost {
	return Ghost{
		Damage:  config.Get().GameConfig.DefaultDamage,
		Symbols: GenerateSymbolsSlice(config.Get().GameConfig.DefaultLenSymbolsSlice),
		X:       config.Get().GameConfig.DefaultLeftPosition - config.Get().GameConfig.DefaultSpriteWidth,
		Speed:   config.Get().GameConfig.DefaultMovementSpeed,
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
	// всего может быть 2 призрака на экране АХТУНГ
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
	if gq.Items[0].Speed > 0 &&
		//gq.Items[0].X >= config.Get().GameConfig.PlayerLeftPosition-config.Get().GameConfig.DefaultSpriteWidth {
		gq.Items[0].X >= config.Get().GameConfig.PlayerLeftPosition {
		hit = true
	} else if gq.Items[0].Speed < 0 &&
		gq.Items[0].X <= config.Get().GameConfig.PlayerRightPosition+config.Get().GameConfig.DefaultSpriteWidth {
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

			score += config.Get().GameConfig.ScoreMatchSymbol
		}

		if len(gq.Items[i].Symbols) != 0 {
			newItems = append(newItems, gq.Items[i])
		} else {
			score += config.Get().GameConfig.ScoreKillGhost
		}
	}

	if len(newItems) != 0 {
		gq.Items = newItems
	}

	log.Println(gq.Items)

	return score
}

func (gq *GhostQueue) Len() int {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	return len(gq.Items)
}

func (gq *GhostQueue) CheckValid() {
	gq.mu.Lock()
	defer gq.mu.Unlock()

	newItems := make([]Ghost, 0, 2)
	for _, i := range gq.Items {
		if len(i.Symbols) != 0 {
			newItems = append(newItems, i)
		}
	}

	gq.Items = newItems
}