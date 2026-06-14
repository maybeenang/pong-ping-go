package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

type PlayerService struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerService {
	return &PlayerService{
		repo: repo,
	}
}

func (s *PlayerService) Register(ctx context.Context, username string) (*domain.Player, error) {
	existing, _ := s.repo.GetByUsername(ctx, username)
	if existing != nil {
		return nil, fmt.Errorf("service: register player : username already exists")
	}

	player := &domain.Player{
		ID:        generateID(),
		Username:  username,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, player); err != nil {
		return nil, fmt.Errorf("service: register player : %w", err)
	}

	return player, nil
}

func (s *PlayerService) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	return s.repo.GetByID(ctx, id)
}

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
