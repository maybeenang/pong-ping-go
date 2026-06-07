package network

import "sync"

type Hub struct {
	mu    sync.RWMutex
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) CreteRoom(name, roomID string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()

	if room, exists := h.Rooms[roomID]; exists {
		return room
	}

	newRoom := NewRoom(name, roomID, h)
	h.Rooms[roomID] = newRoom
	go newRoom.Run()

	return newRoom
}

func (h *Hub) GetRoom(roomID string) *Room {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.Rooms[roomID]
}

func (h *Hub) GetRoomList() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	roomIDs := make([]string, 0, len(h.Rooms))
	for id := range h.Rooms {
		roomIDs = append(roomIDs, id)
	}
	return roomIDs
}

func (h *Hub) RemoveRoom(roomID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.Rooms, roomID)
}
