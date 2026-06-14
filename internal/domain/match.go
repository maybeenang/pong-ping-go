package domain

import "time"

type Match struct {
	ID        string
	RoomID    string
	Player1ID string
	Player2ID string
	WinnerID  string
	Score1    int
	Score2    int
	PlayedAt  time.Time
}
