package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.PlayerRepository = (*PlayerRepo)(nil)

type PlayerRepo struct {
	mu         sync.RWMutex
	players    map[string]*domain.Player
	byUsername map[string]*domain.Player
}

func NewPlayerRepo() *PlayerRepo {
	return &PlayerRepo{
		players:    make(map[string]*domain.Player),
		byUsername: make(map[string]*domain.Player),
	}
}

// Create implements repository.PlayerRepository.
func (p *PlayerRepo) Create(ctx context.Context, player *domain.Player) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.players[player.ID]; exists {
		return fmt.Errorf("player %s already taken", player.ID)
	}

	if _, exists := p.byUsername[player.Username]; exists {
		return fmt.Errorf("player %s already taken", player.Username)
	}

	p.players[player.ID] = player
	p.byUsername[player.Username] = player
	return nil
}

// GetByID implements repository.PlayerRepository.
func (p *PlayerRepo) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	player, exists := p.players[id]
	if !exists {
		return nil, fmt.Errorf("player notfound")

	}

	return player, nil
}

// GetByUsername implements repository.PlayerRepository.
func (p *PlayerRepo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	player, exists := p.byUsername[username]
	if !exists {
		return nil, fmt.Errorf("player notfound")

	}

	return player, nil
}
