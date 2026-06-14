package domain

import "time"

type RoomStatus string

const (
	RoomStatusWaiting  RoomStatus = "waiting"
	RoomStatusPlaying  RoomStatus = "playing"
	RoomStatusFinished RoomStatus = "finished"
)

type Room struct {
	ID        string
	Name      string
	Status    RoomStatus
	CreatedAt time.Time
}
