package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.RoomRepository = (*RoomRepo)(nil)

type RoomRepo struct {
	pool *pgxpool.Pool
}

func NewRoomRepo(pool *pgxpool.Pool) *RoomRepo {
	return &RoomRepo{
		pool: pool,
	}
}

// Create implements repository.RoomRepository.
func (r *RoomRepo) Create(ctx context.Context, room *domain.Room) error {
	query := `
	INSERT INTO rooms (id, name, status, created_at)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.pool.Exec(ctx, query, room.ID, room.Name, room.Status, room.CreatedAt)

	if err != nil {
		return fmt.Errorf("postgree: create room : %w", err)
	}

	return nil
}

// Delete implements repository.RoomRepository.
func (r *RoomRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM rooms
	WHERE id = $1
	`

	_, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("postgres: delete room : %w", err)
	}

	return nil
}

// GetByID implements repository.RoomRepository.
func (r *RoomRepo) GetByID(ctx context.Context, id string) (*domain.Room, error) {
	query := `
	SELECT id, name, status, created_at
	FROM rooms
	WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	room := &domain.Room{}
	var status string

	err := row.Scan(&room.ID, &room.Name, &status, &room.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres: get room by id error : %w", err)
	}

	room.Status = domain.RoomStatus(status)
	return room, nil

}

// List implements repository.RoomRepository.
func (r *RoomRepo) List(ctx context.Context) ([]*domain.Room, error) {
	query := `
	SELECT id, name, status, created_at 
	FROM rooms
	ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("postgres: get list room error: %w", err)
	}
	defer rows.Close()

	var rooms []*domain.Room

	for rows.Next() {
		room := &domain.Room{}
		var status string
		if err := rows.Scan(&room.ID, &room.Name, &status, &room.CreatedAt); err != nil {
			return nil, fmt.Errorf("postgres: scan room: %w", err)
		}
		room.Status = domain.RoomStatus(status)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// UpdateStatus implements repository.RoomRepository.
func (r *RoomRepo) UpdateStatus(ctx context.Context, id string, status domain.RoomStatus) error {
	query := `
	UPDATE rooms 
	SET status = $1
	WHERE id = $2
	`

	result, err := r.pool.Exec(ctx, query, string(status), id)
	if err != nil {
		return fmt.Errorf("postgres: update status room : %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("postgres: room %s not found", id)
	}

	return nil
}
