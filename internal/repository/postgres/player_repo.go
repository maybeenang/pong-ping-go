package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.PlayerRepository = (*PlayerRepo)(nil)

type PlayerRepo struct {
	pool *pgxpool.Pool
}

func NewPlayerRepo(pool *pgxpool.Pool) *PlayerRepo {
	return &PlayerRepo{
		pool: pool,
	}
}

// Create implements repository.PlayerRepository.
func (p *PlayerRepo) Create(ctx context.Context, player *domain.Player) error {
	query := `
	INSERT INTO players (id, username, created_at)
	VALUES ($1, $2, $3)
	`
	_, err := p.pool.Exec(ctx, query, player.ID, player.Username, player.CreatedAt)
	if err != nil {
		return fmt.Errorf("postgres: create player: %w", err)
	}
	return nil
}

// GetByID implements repository.PlayerRepository.
func (p *PlayerRepo) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	query := `
	SELECT id, username, created_at
	FROM players WHERE id = $1
	`

	row := p.pool.QueryRow(ctx, query, id)
	player := &domain.Player{}
	err := row.Scan(&player.ID, &player.Username, &player.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("postgres: get player by id: %w", err)
	}

	return player, nil

}

// GetByUsername implements repository.PlayerRepository.
func (p *PlayerRepo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	query := `
	SELECT id, username, created_at
	FROM players WHERE username = $1
	`

	row := p.pool.QueryRow(ctx, query, username)
	player := &domain.Player{}
	err := row.Scan(&player.ID, &player.Username, &player.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("postgres: get player by username: %w", err)
	}

	return player, nil
}
