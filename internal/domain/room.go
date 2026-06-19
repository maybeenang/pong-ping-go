package domain

import "time"

type RoomStatus string

const (
	RoomStatusWaiting  RoomStatus = "waiting"
	RoomStatusPlaying  RoomStatus = "playing"
	RoomStatusFinished RoomStatus = "finished"
)

type Room struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Status    RoomStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}
