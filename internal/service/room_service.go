// Package service
package service

import (
	"context"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(ctx context.Context, name string) (*domain.Room, error) {
	room := &domain.Room{
		ID:        generateID()[:8],
		Name:      name,
		Status:    domain.RoomStatusWaiting,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) GetRoomByID(ctx context.Context, id string) (*domain.Room, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *RoomService) ListRooms(ctx context.Context) ([]*domain.Room, error) {
	return s.repo.List(ctx)
}

func (s *RoomService) UpdateRoomStatus(ctx context.Context, id string, status domain.RoomStatus) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *RoomService) DeleteRoom(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
