package game

import (
	"fmt"
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

// нужон ли тайм?
type RoomState struct {
	Players     []gameLogic.PlayerCharacter
	Objects     *gameLogic.GhostQueue
	currentTime time.Time
}

type Room struct {
	ID         string
	game       *Game
	MaxPlayers uint
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	dead       chan *Player
	enemyEnd   chan struct{}
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
		dead:       make(chan *Player),
		enemyEnd:   make(chan struct{}),

		mu:     &sync.Mutex{},
		ticker: time.NewTicker(1 * time.Second),
		state: &RoomState{
			Objects: gameLogic.NewGhostStack(),
		},
	}
}

func (r *Room) Run() {
	log.Printf("room loop started ID=%s", r.ID)
	for {
		select {
		case <-r.dead:
			log.Printf("room id=%s, players are dead", r.ID)
			r.Close()

		case <-r.enemyEnd:
			log.Printf("room id=%s, enemy ends", r.ID)
			r.Close()

		case player := <-r.unregister:

			delete(r.Players, player.ID)
			log.Printf("player %s was removed from room", player.ID)

			// убираем вышедшему игроку очки, а оставшемуся очки делим на 2
			for _, players := range r.Players {
				players.SendMessage(&Message{"END", fmt.Sprintf("player %s has left, GAME OVER", player.ID)})
			}

			r.state.Players = r.state.Players[:len(r.state.Players)-1]
			r.playerCnt -= 1

			r.Close()

		case player := <-r.register:
			r.Players[player.ID] = player
			log.Printf("player %s joined", player.ID)
			player.SendMessage(&Message{"CONNECTED", nil})

			r.state.Players = append(r.state.Players, gameLogic.NewPlayerCharacter(player.ID))

			r.playerCnt += 1

			if r.playerCnt == r.MaxPlayers {
				//TODO(): аппендить призраков на каждом тике, но не больше заданного значения
				r.state.Objects.PushBack(
					gameLogic.NewGhost(100, 20, "kek", 10, 5))
				r.state.Objects.PushBack(
					gameLogic.NewGhost(-110, 80, "kek1", 10, 5))
				r.state.Objects.PushBack(
					gameLogic.NewGhost(120, 80, "kek2", 10, 5))
			}

		case <-r.ticker.C:
			if r.playerCnt != r.MaxPlayers {
				continue
			}

			if r.state.Objects.Len() == 0 {
				log.Println("enemy end")
				r.Close()
				//r.enemyEnd <- struct{}{}
				//continue
			}

			log.Println("tick")

			// тут ваша игровая механика
			// взять команды у плеера, обработать их
			r.state.currentTime = time.Now()

			f := r.state.Objects.MoveAllGhosts()
			if f {
				r.state.Objects.PopFront()

				for i := range r.state.Players {
					r.state.Players[i].HP -= 20

					if r.state.Players[i].HP <= 0 {
						r.dead <- &Player{}
						//continue
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
	//player.CloseConn()
	r.unregister <- player
}

func (r *Room) Close() {
	r.mu.Lock()
	for _, player := range r.Players {
		//r.RemovePlayer(player)
		player.SendMessage(&Message{"END", fmt.Sprintf("GAME OVER your score = %d", player.Score)})
		r.game.RemovePlayer(player)
	}
	r.mu.Unlock()
	r.game.CloseRoom(r.ID)
}
