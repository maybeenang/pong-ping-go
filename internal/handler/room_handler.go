package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/service"
)

type RoomHandler struct {
	roomService *service.RoomService
}

func NewRoomHandler(roomService *service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	room := h.roomService.CreateRoom(req.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"room_id": room.ID,
	})
}

func (h *RoomHandler) ListRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{
		"rooms": h.roomService.GetRoomList(),
	})
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.PathValue("id")

	room := h.roomService.GetRoom(roomID)

	if room == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"room": room,
	})
}
