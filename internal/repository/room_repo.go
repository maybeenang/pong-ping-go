// Package repository
package repository

import (
	"context"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
)

type RoomRepository interface {
	Create(ctx context.Context, room *domain.Room) error
	GetByID(ctx context.Context, id string) (*domain.Room, error)
	List(ctx context.Context) ([]*domain.Room, error)
	UpdateStatus(ctx context.Context, id string, status domain.RoomStatus) error
	Delete(ctx context.Context, id string) error
}
