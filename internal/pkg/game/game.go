package game

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"sync"
)

var InstanceGame *Game

func init() {
	// игра крутится как отдельная сущность всегда
	InstanceGame = NewGame(10)
	go panicWorker.PanicWorker(InstanceGame.Run)
}

type Game struct {
	RoomsCount uint32
	mu         *sync.Mutex
	searchMu   *sync.Mutex
	register   chan *Player
	rooms      map[string]*Room
	emptyRooms map[string]*Room
}

func NewGame(roomsCount uint32) *Game {
	return &Game{
		RoomsCount: roomsCount,
		mu:         &sync.Mutex{},
		searchMu:   &sync.Mutex{},
		register:   make(chan *Player, 10),
		rooms:      make(map[string]*Room),
		emptyRooms: make(map[string]*Room),
	}
}

func (g *Game) Run() {
	log.Println("main loop started")

LOOP:
	for player := range g.register {
		//g.searchMu.Lock()
		log.Printf("len empty rooms = %d, len full rooms = %d", len(g.emptyRooms), len(g.rooms))
		for _, room := range g.emptyRooms {

			if len(room.Players) < int(room.MaxPlayers) {
				room.AddPlayer(player)
				g.MakeRoomFull(room)
				continue LOOP
			}
		}

		room := NewRoom(2, g)
		g.AddEmptyRoom(room)
		go panicWorker.PanicWorker(room.Run)

		room.AddPlayer(player)
		//g.searchMu.Unlock()
	}
}

func (g *Game) AddPlayer(player *Player) {
	log.Printf("player %s queued to add", player.Token)
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

func (g *Game) AddEmptyRoom(room *Room) {
	g.mu.Lock()
	g.emptyRooms[room.ID] = room
	g.mu.Unlock()
}

func (g *Game) MakeRoomFull(room *Room) {
	g.mu.Lock()
	g.emptyRooms[room.ID] = nil
	delete(g.emptyRooms, room.ID)

	g.rooms[room.ID] = room
	g.mu.Unlock()
}

func (g *Game) CloseRoom(ID string) {
	g.mu.Lock()
	if _, ok := g.rooms[ID]; !ok {
		log.Println("deleted empty room")
		delete(g.emptyRooms, ID)
	}
	delete(g.rooms, ID)
	g.mu.Unlock()

	log.Printf("room Token=%s deleted", ID)
}
