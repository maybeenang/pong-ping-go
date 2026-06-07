// Package service
package service

import (
	"math/rand"

	"github.com/maybeenang/pong-ping-v2/internal/network"
)

type RoomService struct {
	hub *network.Hub
}

func NewRoomService(hub *network.Hub) *RoomService {
	return &RoomService{
		hub: hub,
	}
}

func generateRoomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	roomID := make([]byte, length)
	for i := range roomID {
		roomID[i] = charset[rand.Intn(len(charset))]
	}
	return string(roomID)
}

func (s *RoomService) CreateRoom(name string) *network.Room {
	roomID := generateRoomID(8)
	return s.hub.CreteRoom(name, roomID)
}

func (s *RoomService) GetRoomList() []string {
	return s.hub.GetRoomList()
}

func (s *RoomService) GetRoom(roomID string) *network.Room {
	return s.hub.GetRoom(roomID)
}
