package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/config"
	"github.com/maybeenang/pong-ping-v2/internal/handler"
	"github.com/maybeenang/pong-ping-v2/internal/network"
	"github.com/maybeenang/pong-ping-v2/internal/repository/memory"
	"github.com/maybeenang/pong-ping-v2/internal/repository/postgres"
	"github.com/maybeenang/pong-ping-v2/internal/service"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	// DB
	pgPool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("postgres: %v", err)
	}

	// repository
	roomRepo := postgres.NewRoomRepo(pgPool)
	playerRepo := postgres.NewPlayerRepo(pgPool)
	matchRepo := postgres.NewMatchRepo(pgPool)
	leaderboardRepo := memory.NewLeaderboardRepo()

	// service
	roomService := service.NewRoomService(roomRepo)
	playerService := service.NewPlayerService(playerRepo)
	matchService := service.NewMatchService(matchRepo, leaderboardRepo)
	leaderboardService := service.NewLeaderboardService(leaderboardRepo, playerRepo)

	// state
	hub := network.NewHub()

	// service

	// handler
	roomHandler := handler.NewRoomHandler(roomService, hub)
	playerHandler := handler.NewPlayerHandler(playerService)
	matchHandler := handler.NewMatchHandler(matchService)
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardService)
	wsHandler := handler.NewWSHandler(hub)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	// rooms
	mux.HandleFunc("POST /api/rooms", roomHandler.CreateRoom)
	mux.HandleFunc("GET /api/rooms", roomHandler.ListRoom)
	mux.HandleFunc("GET /api/rooms/{id}", roomHandler.GetRoom)

	// player
	mux.HandleFunc("POST /api/players", playerHandler.Register)
	mux.HandleFunc("GET /api/players/{id}", playerHandler.GetPlayer)

	// matches
	mux.HandleFunc("GET /api/players/{id}/matches", matchHandler.GetMatchHistory)

	// leaderboard
	mux.HandleFunc("GET /api/leaderboard", leaderboardHandler.GetTop10)
	mux.HandleFunc("GET /api/leaderboard/{id}/rank", leaderboardHandler.GetPlayerRank)

	mux.HandleFunc("/ws", wsHandler.ServeWS)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	go func() {
		log.Printf("server running on port : %v", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("server: shutdown error : %w", err)
	}

	pgPool.Close()

	log.Println("server closed")
}
