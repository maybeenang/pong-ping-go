package repository

import "context"

type LeaderboardEntry struct {
	PlayerID string
	Wins     int
	Rank     int64
}

type LeaderboardRepository interface {
	IncrementWin(ctx context.Context, playerID string) error
	GetTopN(ctx context.Context, n int) ([]LeaderboardEntry, error)
	GetRank(ctx context.Context, playerID string) (int64, error)
}
