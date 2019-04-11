package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gameLogic"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
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
	Players []PlayerState
	// TODO(): сделать стак (первый призрак, всегда ближе к плееру)
	Objects     *gameLogic.GhostQueue
	CurrentTime time.Time
}

type Room struct {
	ID         string
	game       *Game
	MaxPlayers uint
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState
	playerCnt  uint
}

func NewRoom(maxPlayers uint, game *Game) *Room {
	return &Room{
		ID:         session.GenerateToken(),
		game:       game,
		MaxPlayers: maxPlayers,
		Players:    make(map[string]*Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		mu:         &sync.Mutex{},
		ticker:     time.NewTicker(1 * time.Second),
		state: &RoomState{
			Objects: gameLogic.NewGhostStack(),
		},
	}
}

func (r *Room) Run() {
	log.Printf("room loop started ID=%s", r.ID)
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
				//TODO(): аппендить призраков на каждом тике, но не больше заданного значения
				r.state.Objects.PushBack(
					gameLogic.NewGhost(100, 20, "kek", 10, 5))
			}

		case <-r.ticker.C:
			if r.playerCnt != r.MaxPlayers {
				continue
			}

			if r.state.Objects.Len() == 0 {
				r.Close()
			}

			log.Println("tick")

			// тут ваша игровая механика
			// взять команды у плеера, обработать их
			r.state.CurrentTime = time.Now()

			f := r.state.Objects.MoveAllGhosts()
			if f {
				r.state.Objects.PopFront()

				for i := range r.state.Players {
					r.state.Players[i].HP -= 20
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
	//player.CloseConn()
	r.unregister <- player
}

func (r *Room) Close() {
	r.mu.Lock()
	for _, player := range r.Players {
		//r.RemovePlayer(player)
		r.game.RemovePlayer(player)
	}
	r.mu.Unlock()
	r.game.CloseRoom(r.ID)
}
