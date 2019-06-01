package game

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/stats"
	grpcAuth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"sync"
)

var InstanceGame *Game

func Start(roomsCount uint32, authGRPCConn grpcAuth.AuthCheckerClient) {
	// игра крутится как отдельная сущность всегда
	InstanceGame = NewGame(roomsCount, authGRPCConn)
	// упала игра -> дочерние горутины должны упасть, как и весь сервис
	//go panicWorker.PanicWorker(InstanceGame.Run)
	go InstanceGame.Run()

	log.Println("InstanceGame.Run()")
}

type PlayerWithLink struct {
	Player *Player
	RoomID string
}

// register - просто игроки
// registerUnique - игроки, создающие выделенные комнаты
// registerUniqueLink - игроки, пытающиеся законнектиться к выделенной комнате
type Game struct {
	GRPC               grpcAuth.AuthCheckerClient
	RoomsCount         uint32
	mu                 *sync.Mutex
	searchMu           *sync.Mutex
	register           chan *Player
	registerUnique     chan *Player
	registerUniqueLink chan *PlayerWithLink
	rooms              map[string]*Room
	emptyRooms         map[string]*Room
	emptyRoomsUnique   map[string]*Room
}

func NewGame(roomsCount uint32, authGRPCConn grpcAuth.AuthCheckerClient) *Game {
	return &Game{
		GRPC:               authGRPCConn,
		RoomsCount:         roomsCount,
		mu:                 &sync.Mutex{},
		searchMu:           &sync.Mutex{},
		register:           make(chan *Player, 10),
		registerUnique:     make(chan *Player, 10),
		registerUniqueLink: make(chan *PlayerWithLink, 10),
		rooms:              make(map[string]*Room),
		emptyRooms:         make(map[string]*Room),
		emptyRoomsUnique:   make(map[string]*Room),
	}
}

func (g *Game) Run() {
	log.Println("main loop started")

LOOP:
	for {
		select {
		case player := <-g.register:
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

		case player := <-g.registerUnique:
			log.Println("adding user to unique room")
			room := NewRoom(2, g)
			g.AddUniqueRoom(room)
			go panicWorker.PanicWorker(room.Run)

			room.AddPlayer(player)
			player.SendMessage(&Message{
				Type:    MessageLink,
				Payload: fmt.Sprintf("https://5factorial.tech/%s", room.ID),
			})

		case playerWithLink := <-g.registerUniqueLink:
			log.Printf("adding user by link to unique room = %d", playerWithLink.RoomID)
			if room, ok := g.emptyRoomsUnique[playerWithLink.RoomID]; ok {
				room.AddPlayer(playerWithLink.Player)
				g.MakeUniqueRoomFull(room)
			}
		}
	}
}

func (g *Game) AddPlayer(player *Player) {
	log.Printf("player %s queued to add", player.Token)
	g.register <- player
}

func (g *Game) AddUniquePlayer(player *Player) {
	log.Printf("player %s queued to add, unique room", player.Token)
	g.registerUnique <- player
}

func (g *Game) AddUniquePlayerLink(player *Player, roomID string) {
	log.Printf("player %s queued to add, unique room by link %s", player.Token, roomID)
	g.registerUniqueLink <- &PlayerWithLink{
		Player: player,
		RoomID: roomID,
	}
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
	stats.Stats.AddActiveRoom()
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
	stats.Stats.RemoveActiveRoom()
	if _, ok := g.rooms[ID]; !ok {
		log.Println("deleted empty room")
		delete(g.emptyRooms, ID)
	}
	delete(g.rooms, ID)
	g.mu.Unlock()

	log.Printf("room Token=%s deleted", ID)
}

func (g *Game) AddUniqueRoom(room *Room) {
	g.mu.Lock()
	defer g.mu.Unlock()

	stats.Stats.AddActiveRoom()
	g.emptyRoomsUnique[room.ID] = room
}

func (g *Game) MakeUniqueRoomFull(room *Room) {
	g.mu.Lock()
	g.emptyRoomsUnique[room.ID] = nil
	delete(g.emptyRoomsUnique, room.ID)

	g.rooms[room.ID] = room
	g.mu.Unlock()
}

func (g *Game) CheckRoomLink(roomID string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	_, ok := g.emptyRoomsUnique[roomID]
	return ok
}
