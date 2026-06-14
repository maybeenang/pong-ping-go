package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/service"
)

type LeaderboardHandler struct {
	service *service.LeaderboardService
}

func NewLeaderboardHandler(service *service.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{
		service: service,
	}
}

func (h *LeaderboardHandler) GetTop10(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.GetTop10(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"entries": entries,
	})
}

func (h *LeaderboardHandler) GetPlayerRank(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	rank, err := h.service.GetPlayerRank(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"player_id": id,
		"rank":      rank,
	})
}
