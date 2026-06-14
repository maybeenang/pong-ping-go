package repository

import (
	"context"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) error
	GetByID(ctx context.Context, id string) (*domain.Player, error)
	GetByUsername(ctx context.Context, username string) (*domain.Player, error)
}
