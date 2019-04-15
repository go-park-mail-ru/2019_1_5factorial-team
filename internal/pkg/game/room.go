package game

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gameLogic"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
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

	playersInputs []gameLogic.Symbol
	playerInput   chan *gameLogic.Symbol
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
		ticker: time.NewTicker(2 * time.Second),
		state: &RoomState{
			Objects: gameLogic.NewGhostStack(),
		},

		playerInput:   make(chan *gameLogic.Symbol),
		playersInputs: make([]gameLogic.Symbol, 0, 10),
	}
}

func (r *Room) Run() {
	log.Printf("room loop started Token=%s", r.ID)
	//LOOP:
	for {
		select {
		case in := <-r.playerInput:
			r.playersInputs = append(r.playersInputs, *in)

		case <-r.dead:
			log.Printf("room id=%s, players are dead", r.ID)
			r.Close()
			return

		case <-r.enemyEnd:
			log.Printf("room id=%s, enemy ends", r.ID)
			r.Close()
			return

		case player := <-r.unregister:
			delete(r.Players, player.Token)
			log.Printf("player %s was removed from room", player.Token)

			// убираем вышедшему игроку очки, а оставшемуся очки делим на 2
			for _, players := range r.Players {
				players.SendMessage(&Message{"END",
					fmt.Sprintf("player %s has left, GAME OVER", player.Token)})
			}

			r.state.Players = r.state.Players[:len(r.state.Players)-1]
			r.playerCnt -= 1

			r.Close()
			return

		case player := <-r.register:
			r.Players[player.Token] = player
			log.Printf("player %s joined", player.Token)
			player.SendMessage(&Message{"CONNECTED", nil})

			r.state.Players = append(r.state.Players, gameLogic.NewPlayerCharacter(player.Token))

			r.playerCnt += 1

		case <-r.ticker.C:
			if r.playerCnt != r.MaxPlayers {
				continue
			}

			if len(r.playersInputs) != 0 {
				for _, val := range r.playersInputs {
					r.state.Objects.PopSymbol(val)
				}
				r.playersInputs = make([]gameLogic.Symbol, 0, 10)
			}

			if r.state.Objects.Len() <= 4 {
				r.state.Objects.PushBack(gameLogic.NewRandomGhost())
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
						log.Println("---===DEAD===---")
						r.Close()
						return

						// на каналах не работает хз поч
						//r.dead <- &Player{}
						//continue LOOP
					}
				}
			}

			for _, player := range r.Players {
				player.SendState(r.state)
			}

			fmt.Println(r.playersInputs)

			//r.PrintStates()
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

func (r *Room) PrintStates() {
	log.Printf("================= room id = %s =================", r.ID)

	//log.Printf("players:")
	//for i, val := range r.Players {
	//	log.Printf("\t%s: %v", i, val)
	//}

	log.Printf("state:\t\nplayers:")
	for i, val := range r.state.Players {
		log.Printf("\t\t%d: %v", i, val)
	}
	log.Printf("\t\nghosts:\n%v", r.state.Objects)

	log.Printf("================= =================")
}
