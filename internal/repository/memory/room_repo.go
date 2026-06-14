// Package memory
package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.RoomRepository = (*RoomRepo)(nil)

type RoomRepo struct {
	mu    sync.RWMutex
	rooms map[string]*domain.Room
}

func NewRoomRepo() *RoomRepo {
	return &RoomRepo{
		rooms: make(map[string]*domain.Room),
	}
}

// Create implements repository.RoomRepository.
func (r *RoomRepo) Create(ctx context.Context, room *domain.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.rooms[room.ID]; exists {
		return fmt.Errorf("room %s already exists", room.ID)
	}

	r.rooms[room.ID] = room
	return nil
}

// Delete implements repository.RoomRepository.
func (r *RoomRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.rooms[id]; !exists {
		return fmt.Errorf("room %s not found", id)
	}

	delete(r.rooms, id)
	return nil
}

// GetByID implements repository.RoomRepository.
func (r *RoomRepo) GetByID(ctx context.Context, id string) (*domain.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, exists := r.rooms[id]
	if !exists {
		return nil, fmt.Errorf("room %s not found", id)
	}

	return room, nil
}

// List implements repository.RoomRepository.
func (r *RoomRepo) List(ctx context.Context) ([]*domain.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*domain.Room, 0, len(r.rooms))

	for _, room := range r.rooms {
		result = append(result, room)
	}

	return result, nil

}

// UpdateStatus implements repository.RoomRepository.
func (r *RoomRepo) UpdateStatus(ctx context.Context, id string, status domain.RoomStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	room, exists := r.rooms[id]
	if !exists {
		return fmt.Errorf("room %s not exists", id)
	}

	room.Status = status
	return nil
}
