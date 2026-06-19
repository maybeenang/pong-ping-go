package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maybeenang/pong-ping-v2/internal/config"
	"github.com/maybeenang/pong-ping-v2/internal/handler"
	"github.com/maybeenang/pong-ping-v2/internal/middleware"
	"github.com/maybeenang/pong-ping-v2/internal/network"
	"github.com/maybeenang/pong-ping-v2/internal/repository/memory"
	"github.com/maybeenang/pong-ping-v2/internal/repository/postgres"
	"github.com/maybeenang/pong-ping-v2/internal/service"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	jsonLogger := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonLogger)

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

	handler := middleware.Logger(logger)(mux)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
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
		log.Printf("server: shutdown error : %v", err)
	}

	pgPool.Close()

	log.Println("server closed")
}
