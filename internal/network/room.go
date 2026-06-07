package network

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/game"
)

type RegisterRequest struct {
	Client   *Client
	Response chan registerResponse
}

type registerResponse struct {
	PlayerID int
	Err      error
}

type Room struct {
	ID     string
	Name   string
	Engine *game.Engine
	Hub    *Hub

	clients map[*Client]bool

	Register   chan RegisterRequest
	Unregister chan *Client
}

func NewRoom(name, id string, hub *Hub) *Room {
	return &Room{
		ID:     id,
		Engine: game.NewEngine(),
		Hub:    hub,
		Name:   name,

		clients: make(map[*Client]bool),

		Register:   make(chan RegisterRequest),
		Unregister: make(chan *Client),
	}
}

func safeClose(ch chan []byte) {
	defer func() { recover() }()
	close(ch)
}

func (r *Room) Run() {
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	log.Printf("Room running with id: %s\n", r.ID)

	for {
		select {

		case req := <-r.Register:
			if len(r.clients) >= 2 {
				req.Response <- registerResponse{Err: fmt.Errorf("Room Full")}
				continue
			}
			playerID := len(r.clients) + 1
			req.Client.PlayerID = playerID
			r.clients[req.Client] = true
			log.Printf("player %d joined room %s", playerID, r.ID)
			req.Response <- registerResponse{PlayerID: playerID}

		case client := <-r.Unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				safeClose(client.Send)
				log.Printf("player %d left room %s", client.PlayerID, r.ID)
			}

			if len(r.clients) == 0 {
				log.Printf("room %s empty, removing", r.ID)
				r.Hub.RemoveRoom(r.ID)
				return
			}

		case <-ticker.C:
			if len(r.clients) == 2 {
				r.Engine.Update()

				stateBytes, err := json.Marshal(r.Engine.State)
				if err != nil {
					continue
				}

				// broadcast langsung di goroutine ini — tidak perlu goroutine terpisah
				for client := range r.clients {
					select {
					case client.Send <- stateBytes:
					default:
						// client lambat: putus koneksi
						delete(r.clients, client)
						safeClose(client.Send)
						log.Printf("player %d slow, kicked from room %s", client.PlayerID, r.ID)
					}
				}
			}
		}
	}
}
