package service

import (
	"context"
	"sync"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/domain"
	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

type MatchService struct {
	wg              sync.WaitGroup
	matchRepo       repository.MatchRepository
	leaderboardRepo repository.LeaderboardRepository
}

func NewMatchService(
	matchRepo repository.MatchRepository,
	leaderboardRepo repository.LeaderboardRepository,
) *MatchService {
	return &MatchService{
		matchRepo:       matchRepo,
		leaderboardRepo: leaderboardRepo,
	}
}

func (s *MatchService) RecordMatchResult(
	ctx context.Context,
	player1ID, player2ID string,
	player1Score, player2Score int,
	roomID string,
) (*domain.Match, error) {
	winnerID := player1ID
	if player1Score < player2Score {
		winnerID = player2ID
	}

	match := &domain.Match{
		ID:        generateID(),
		RoomID:    roomID,
		Player1ID: player1ID,
		Player2ID: player2ID,
		WinnerID:  winnerID,
		Score1:    player1Score,
		Score2:    player2Score,
		PlayedAt:  time.Now(),
	}

	if err := s.matchRepo.Create(ctx, match); err != nil {
		return nil, err
	}

	s.wg.Add(1)
	asyncCtx := context.WithoutCancel(ctx)
	asyncCtx, cancel := context.WithTimeout(asyncCtx, 5*time.Second)
	defer cancel()

	go func() {
		defer s.wg.Done()
		s.leaderboardRepo.IncrementWin(asyncCtx, winnerID)
	}()

	return match, nil
}

func (s *MatchService) GetMatchHistory(ctx context.Context, playerID string) ([]*domain.Match, error) {
	return s.matchRepo.GetByPlayerID(ctx, playerID)
}
