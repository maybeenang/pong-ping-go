package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/service"
)

type MatchHandler struct {
	service *service.MatchService
}

func NewMatchHandler(s *service.MatchService) *MatchHandler {
	return &MatchHandler{service: s}
}

func (h *MatchHandler) GetMatchHistory(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	mathces, err := h.service.GetMatchHistory(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"matches": mathces,
	})
}
