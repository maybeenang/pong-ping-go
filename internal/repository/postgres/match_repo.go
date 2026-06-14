package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.MatchRepository = (*MatchRepo)(nil)

type MatchRepo struct {
	pool *pgxpool.Pool
}

func NewMatchRepo(pool *pgxpool.Pool) *MatchRepo {
	return &MatchRepo{
		pool: pool,
	}
}

// Create implements repository.MatchRepository.
func (m *MatchRepo) Create(ctx context.Context, match *domain.Match) error {
	query := `
	INSERT INTO matches (id, room_id, player1_id, player2_id, winner_id, score_1, score_2, played_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := m.pool.Exec(ctx, query,
		match.ID,
		match.RoomID,
		match.Player1ID,
		match.Player2ID,
		match.WinnerID,
		match.Score1,
		match.Score2,
		match.PlayedAt,
	)

	if err != nil {
		return fmt.Errorf("postgres: create match : %w", err)
	}

	return nil
}

// GetByPlayerID implements repository.MatchRepository.
func (m *MatchRepo) GetByPlayerID(ctx context.Context, playerID string) ([]*domain.Match, error) {
	query := `
	SELECT id, room_id, player1_id, player2_id, winner_id, score_1, score_2, played_at
	FROM matches
	WHERE player1_id = $1 OR player2_id = $1
	ORDER BY played_at DESC
	`

	return m.queryMatches(ctx, query, playerID)
}

// GetByRoomID implements repository.MatchRepository.
func (m *MatchRepo) GetByRoomID(ctx context.Context, roomID string) ([]*domain.Match, error) {
	query := `
	SELECT id, room_id, player1_id, player2_id, winner_id, score_1, score_2, played_at
	FROM matches
	WHERE room_id = $1
	ORDER BY played_at DESC
	`

	return m.queryMatches(ctx, query, roomID)
}

func (m *MatchRepo) queryMatches(ctx context.Context, query string, arg any) ([]*domain.Match, error) {
	rows, err := m.pool.Query(ctx, query, arg)
	if err != nil {
		return nil, fmt.Errorf("posgres: error query matches : %w", err)
	}
	defer rows.Close()

	var matches []*domain.Match

	for rows.Next() {
		match := &domain.Match{}
		if err := rows.Scan(
			&match.ID,
			&match.RoomID,
			&match.Player1ID,
			&match.Player2ID,
			&match.WinnerID,
			&match.Score1,
			&match.Score2,
			&match.PlayedAt,
		); err != nil {
			return nil, fmt.Errorf("postgres: error scanning match : %w", err)
		}

		matches = append(matches, match)

	}
	return matches, nil
}

