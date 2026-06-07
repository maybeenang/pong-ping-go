package main

import (
	"fmt"
	"net/http"

	"github.com/maybeenang/pong-ping-v2/internal/handler"
	"github.com/maybeenang/pong-ping-v2/internal/network"
	"github.com/maybeenang/pong-ping-v2/internal/service"
)

func main() {

	// state
	hub := network.NewHub()

	// service
	roomService := service.NewRoomService(hub)

	// handler
	roomHandler := handler.NewRoomHandler(roomService)
	wsHandler := handler.NewWSHandler(hub)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	mux.HandleFunc("/api/create-room", roomHandler.CreateRoom)
	mux.HandleFunc("/api/rooms", roomHandler.ListRoom)

	mux.HandleFunc("/api/rooms/{id}", roomHandler.GetRoom)

	mux.HandleFunc("/ws", wsHandler.ServeWS)

	port := ":8080"
	fmt.Printf("Server running on localhost:%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
