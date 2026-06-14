package memory

import (
	"context"
	"sort"
	"sync"

	"github.com/maybeenang/pong-ping-v2/internal/repository"
)

var _ repository.LeaderboardRepository = (*LeaderbordRepo)(nil)

type LeaderbordRepo struct {
	mu   sync.RWMutex
	wins map[string]int64
}

func NewLeaderboardRepo() *LeaderbordRepo {
	return &LeaderbordRepo{
		wins: make(map[string]int64),
	}
}

// GetRank implements repository.LeaderboardRepository.
func (l *LeaderbordRepo) GetRank(ctx context.Context, playerID string) (int64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	wins, ok := l.wins[playerID]
	if !ok {
		return 0, nil
	}

	rank := int64(1)
	for _, w := range l.wins {
		if w > wins {
			rank++
		}
	}

	return rank, nil
}

// GetTopN implements repository.LeaderboardRepository.
func (l *LeaderbordRepo) GetTopN(ctx context.Context, n int) ([]repository.LeaderboardEntry, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	entries := make([]repository.LeaderboardEntry, 0, len(l.wins))

	for playerID, wins := range l.wins {
		entries = append(entries, repository.LeaderboardEntry{
			PlayerID: playerID,
			Wins:     int(wins),
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Wins > entries[j].Wins
	})

	if n < len(entries) {
		entries = entries[:n]
	}

	for i := range entries {
		entries[i].Rank = int64(i + 1)
	}

	return entries, nil
}

func (l *LeaderbordRepo) IncrementWin(ctx context.Context, playerID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.wins[playerID]++
	return nil
}
