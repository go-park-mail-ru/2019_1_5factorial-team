package game

import (
	"log"
	"sync"
	"time"
)

// ось X (-100; 100) => игроки стоят в нуле

// X - середина
type PlayerState struct {
	ID string
	X  int
	HP int
}

// X - середина
type ObjectState struct {
	ID    string
	Type  string
	X     int
	Speed int
}

type RoomState struct {
	Players     []PlayerState
	Objects     []ObjectState
	CurrentTime time.Time
}

type Room struct {
	ID         string
	MaxPlayers uint
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState
	playerCnt  uint
}

func NewRoom(maxPlayers uint) *Room {
	return &Room{
		MaxPlayers: maxPlayers,
		Players:    make(map[string]*Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(1 * time.Second),
		state:      &RoomState{},
	}
}

func (r *Room) Run() {
	log.Println("room loop started")
	for {
		select {
		case player := <-r.unregister:
			delete(r.Players, player.ID)
			log.Printf("player %s was removed from room", player.ID)

			r.state.Players = r.state.Players[:len(r.state.Players)-1]
			r.playerCnt -= 1
		case player := <-r.register:
			r.Players[player.ID] = player
			log.Printf("player %s joined", player.ID)
			player.SendMessage(&Message{"CONNECTED", nil})

			r.state.Players = append(r.state.Players, PlayerState{player.ID, 0, 100})

			r.playerCnt += 1

			if r.playerCnt == r.MaxPlayers {
				r.state.Objects = append(r.state.Objects, ObjectState{ID: "kek", Type: "gh", X: 100, Speed: -10})
			}

		case <-r.ticker.C:
			if r.playerCnt != r.MaxPlayers {
				continue
			}

			log.Println("tick")

			// тут ваша игровая механика
			// взять команды у плеера, обработать их
			r.state.CurrentTime = time.Now()
			for i := range r.state.Objects {
				r.state.Objects[i].X += r.state.Objects[i].Speed
				if r.state.Objects[i].X == 0 {
					for i := range r.state.Players {
						r.state.Players[i].HP -= 20
					}
				}
			}

			for _, player := range r.Players {
				player.SendState(r.state)
			}
		}
	}
}

func (r *Room) AddPlayer(player *Player) {
	player.room = r
	r.register <- player
}

func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}
