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
	rooms      map[string]*Room
}

func NewGame(roomsCount uint32) *Game {
	return &Game{
		RoomsCount: roomsCount,
		mu:         &sync.Mutex{},
		register:   make(chan *Player),
		rooms:      make(map[string]*Room),
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

		room := NewRoom(2, g)
		g.AddRoom(room)
		go room.Run()

		room.AddPlayer(player)
	}
}

func (g *Game) AddPlayer(player *Player) {
	log.Printf("player %s queued to add", player.ID)
	g.register <- player
}

func (g *Game) RemovePlayer(player *Player) {
	player.CloseConn()
}

func (g *Game) AddRoom(room *Room) {
	g.mu.Lock()
	g.rooms[room.ID] = room
	g.mu.Unlock()
}

func (g *Game) CloseRoom(ID string) {
	g.mu.Lock()
	delete(g.rooms, ID)
	g.mu.Unlock()

	log.Printf("room ID=%s deleted", ID)
}
