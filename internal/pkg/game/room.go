package game

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gameLogic"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

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
	gameTime   float64
	pause      bool

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
		pause:      false,
		mu:         &sync.Mutex{},
		ticker:     time.NewTicker(config.Get().GameConfig.TickerTime.Duration),
		state: &RoomState{
			Objects: gameLogic.NewGhostStack(),
		},

		playerInput:   make(chan *gameLogic.Symbol),
		playersInputs: make([]gameLogic.Symbol, 0, 10),
	}
}

func (r *Room) Run() {
	log.Printf("room loop started Token=%s", r.ID)

	for {
		select {
		case in := <-r.playerInput:
			if r.playerCnt != r.MaxPlayers || r.pause {
				log.Println("skip player input, bcs game not started, waiting second player OR room paused")
				for _, player := range r.Players {
					player.SendMessage(&Message{
						Type:    MessagePause,
						Payload: nil,
					})
				}
				continue
			}
			// TODO(): если игра не началась или закончилась, но юзер шлет мувы, то они записываются
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
			r.endGame(player)
			return

		case player := <-r.register:
			r.addPlayerToState(player)

		case <-r.ticker.C:
			if r.playerCnt != r.MaxPlayers || r.pause {
				continue
			}

			r.mu.Lock()

			r.updateRoomState()

			for _, player := range r.Players {
				player.SendState(r.state)
			}

			r.mu.Unlock()
		}
	}
}

func (r *Room) endGame(player *Player) {
	delete(r.Players, player.Token)
	log.Printf("player %s was removed from room", player.Token)

	pc := ""
	if r.state.Players[0].Token == player.Token {
		pc = r.state.Players[0].Nick
	} else {
		pc = r.state.Players[1].Nick
	}

	// TODO(): убираем вышедшему игроку очки, а оставшемуся очки делим на 2
	// TODO(): добавить баланс (хочу очки деленные на 100 прибавлять к балансу)
	// TODO(): записывать в борду максимальный счет

	for _, players := range r.Players {
		players.SendMessage(&Message{
			Type:    MessageEnd,
			Payload: fmt.Sprintf("player %s has left, GAME OVER", pc),
		})
	}

	r.state.Players = r.state.Players[:len(r.state.Players)-1]
	r.playerCnt -= 1

	r.Close()
}

func (r *Room) addPlayerToState(player *Player) {
	r.Players[player.Token] = player

	log.Printf("player %s joined", player.Token)

	player.SendMessage(&Message{
		Type:    MessageConnect,
		Payload: nil,
	})

	npc, err := gameLogic.NewPlayerCharacter(player.Token, r.game.GRPC)
	if err != nil {
		log.Error(err.Error(), "cant create character in room, closing")

		r.Close()
		return
	}

	r.state.Players = append(r.state.Players, npc)

	r.playerCnt += 1
}

func (r *Room) rakePlayerInputs() {
	log.Println("rake", r.state.Players[0].HP, r.state.Players[1].HP)

	log.Println("===")
	for _, val := range r.playersInputs {
		val := r.state.Objects.PopSymbol(val)

		for i := range r.state.Players {
			r.state.Players[i].Score += val
		}
	}
	log.Println("===")

	r.playersInputs = make([]gameLogic.Symbol, 0, 10)
}

func (r *Room) updateRoomState() {
	//r.state.currentTime = time.Now()
	// TODO(): тут типа прибавляем тик в мкс
	r.gameTime += 1000000

	if len(r.playersInputs) != 0 {
		r.rakePlayerInputs()
	}

	r.state.Objects.CheckValid()

	if r.state.Objects.Len() < 2 && rand.Float64() < 1-math.Pow(0.993, r.gameTime) {
		// игровая логика
		//r.state.Objects.PushBack(gameLogic.NewRandomGhost())
		r.state.Objects.AddNewGhost()
		log.Println("added new ghost")
		log.Println(r.state.Objects)
	}

	f := r.state.Objects.MoveAllGhosts()
	if f {
		deletedGhost := r.state.Objects.PopFront()

		if deletedGhost.Speed > 0 {
			r.state.Players[0].HP -= int(deletedGhost.Damage)

		} else if deletedGhost.Speed < 0 {
			r.state.Players[1].HP -= int(deletedGhost.Damage)

			//if r.state.Players[1].HP <= 0 {
			//	r.Close()
			//	return
			//
			//	// на каналах не работает хз поч
			//	//r.dead <- &Player{}
			//	//continue LOOP
			//}
		}

		if r.state.Players[0].HP <= 0 || r.state.Players[1].HP <= 0 {
			r.Close()

			return
		}
	}

	//log.Println("tick")
}

func (r *Room) SendMessageAllPlayers() {

}

func (r *Room) AddPlayer(player *Player) {
	player.room = r
	r.register <- player
}

func (r *Room) RemovePlayer(player *Player) {
	//player.CloseConn()
	r.unregister <- player
}

func (r *Room) Pause() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.pause = true
}

func (r *Room) Resume() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.pause = false
}

func (r *Room) Close() {
	r.mu.Lock()

	r.ticker.Stop()

	ctx := context.Background()

	// добавляю юзерам очки их персонажей
	for _, p := range r.state.Players {
		if _, ok := r.Players[p.Token]; ok {
			r.Players[p.Token].Score = p.Score

			uID, err := r.game.GRPC.GetIDFromSession(ctx, &grpcAuth.Cookie{Token: p.Token, Expiration: ""})
			if err != nil {
				log.Error(errors.Wrap(err, "cant create user, GetID"))
				r.Players[p.Token].Score = 0
				continue
			}

			_, err = r.game.GRPC.UpdateScore(ctx, &grpcAuth.UpdateScoreReq{ID: uID.ID, Score: int32(p.Score)})
			if err != nil {
				log.Error("cant update user score, user id=%s, token=%s, score=%d, err=%s",
					uID.ID, p.Token, p.Score, err.Error())

				r.Players[p.Token].Score = 0
				continue
			}
		}
	}

	for _, player := range r.Players {
		player.SendMessage(&Message{
			Type:    MessageEnd,
			Payload: fmt.Sprintf("GAME OVER your score = %d", player.Score),
		})
		// по идеи убирать игрока надо здесь, а не через game
		r.game.RemovePlayer(player)
	}

	r.mu.Unlock()
	r.game.CloseRoom(r.ID)
}
