package service

import (
	"context"

	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

type LeaderboardService struct {
	leaderboardRepo repository.LeaderboardRepository
	playerRepo      repository.PlayerRepository
}

func NewLeaderboardService(lRepo repository.LeaderboardRepository, pRepo repository.PlayerRepository) *LeaderboardService {
	return &LeaderboardService{
		leaderboardRepo: lRepo,
		playerRepo:      pRepo,
	}
}

func (s *LeaderboardService) GetTop10(ctx context.Context) ([]repository.LeaderboardEntry, error) {
	entries, err := s.leaderboardRepo.GetTopN(ctx, 10)
	if err != nil {
		return nil, err
	}

	for i := range entries {
		player, err := s.playerRepo.GetByID(ctx, entries[i].PlayerID)
		if err != nil {
			continue
		}

		entries[i].Username = player.Username
	}

	return entries, nil
}

func (s *LeaderboardService) GetPlayerRank(ctx context.Context, playerID string) (int64, error) {
	return s.leaderboardRepo.GetRank(ctx, playerID)
}
