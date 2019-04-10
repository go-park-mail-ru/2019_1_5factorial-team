package game

import (
	"log"
	"sync"
)

var InstanceGame *Game

func init() {
	// игра крутится как отдельная сущность всегда
	InstanceGame = NewGame(10)
	go InstanceGame.Run()
}

type Game struct {
	RoomsCount uint32
	mu         *sync.Mutex
	register   chan *Player
	rooms      []*Room
}

func NewGame(roomsCount uint32) *Game {
	return &Game{
		RoomsCount: roomsCount,
		mu: &sync.Mutex{},
		register: make(chan *Player),
	}
}

func (g *Game) Run() {
	log.Println("main loop started")

LOOP:
	for {
		player := <-g.register

		for _, room := range g.rooms {
			if len(room.Players) < int(room.MaxPlayers) {
				room.AddPlayer(player)
				continue LOOP
			}
		}

		room := NewRoom(2)
		g.AddRoom(room)
		go room.Run()

		room.AddPlayer(player)
	}
}

func (g *Game) AddPlayer(player *Player) {
	log.Printf("player %s queued to add", player.ID)
	g.register <- player
}

func (g *Game) AddRoom(room *Room) {
	g.rooms = append(g.rooms, room)
}
