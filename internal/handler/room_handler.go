package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/network"
	"github.com/maybeenang/pong-ping-v2/internal/service"
)

type RoomHandler struct {
	roomService *service.RoomService
	hub         *network.Hub
}

func NewRoomHandler(roomService *service.RoomService, hub *network.Hub) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
		hub:         hub,
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

	room, err := h.roomService.CreateRoom(r.Context(), req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.hub.CreteRoom(room.Name, room.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"room_id":     room.ID,
		"room_name":   room.Name,
		"room_status": string(room.Status),
	})
}

func (h *RoomHandler) ListRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rooms, err := h.roomService.ListRooms(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"rooms": rooms,
	})
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.PathValue("id")

	if roomID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.GetRoomByID(r.Context(), roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if room == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"room": room,
	})
}
