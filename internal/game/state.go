// Package game
package game

type GameState struct {
	BallX   float64 `json:"ball_x"`
	BallY   float64 `json:"ball_y"`
	Paddle1 float64 `json:"paddle_1"`
	Paddle2 float64 `json:"paddle_2"`
}

func NewGameState() *GameState {
	return &GameState{
		BallX:   50,
		BallY:   50,
		Paddle1: 50,
		Paddle2: 50,
	}
}
