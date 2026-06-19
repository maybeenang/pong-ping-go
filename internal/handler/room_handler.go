package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/network"
	"github.com/maybeenang/pong-ping-v2/internal/service"
	"github.com/maybeenang/pong-ping-v2/pkg/response"
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
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	room, err := h.roomService.CreateRoom(r.Context(), req.Name)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	h.hub.CreteRoom(room.Name, room.ID)

	response.Success(w, http.StatusCreated, "Create room success", map[string]any{
		"room": room,
	})

}

func (h *RoomHandler) ListRoom(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.roomService.ListRooms(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.Success(w, http.StatusOK, "Get data success", map[string]any{
		"rooms": rooms,
	})
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.PathValue("id")

	if roomID == "" {
		response.Error(w, http.StatusBadRequest, errors.New("room ID is required"))
		return
	}

	room, err := h.roomService.GetRoomByID(r.Context(), roomID)
	if err != nil {
		response.Error(w, http.StatusNotFound, errors.New("room not found"))
		return
	}

	response.Success(w, http.StatusOK, "Get room by id success", map[string]any{
		"room": room,
	})
}
