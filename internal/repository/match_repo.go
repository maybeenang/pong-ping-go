package repository

import (
	"context"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
)

type MatchRepository interface {
	Create(ctx context.Context, match *domain.Match) error
	GetByRoomID(ctx context.Context, roomID string) ([]*domain.Match, error)
	GetByPlayerID(ctx context.Context, playerID string) ([]*domain.Match, error)
}
