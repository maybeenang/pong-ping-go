package game

import "sync"

type Engine struct {
	State     *GameState
	VelocityX float64
	VelocityY float64
	mu        sync.Mutex
}

func NewEngine() *Engine {
	return &Engine{
		State:     NewGameState(),
		VelocityX: 1,
		VelocityY: 1,
	}
}

func (e *Engine) Update() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.State.BallX += e.VelocityX
	e.State.BallY += e.VelocityY

	if e.State.BallY <= 0 || e.State.BallY >= 100 {
		e.VelocityY = -e.VelocityY
	}

	halfPaddle := 10.0

	// check if ball in left
	if e.State.BallX <= 3 {
		if e.State.BallY >= (e.State.Paddle1-halfPaddle) && e.State.BallY <= (e.State.Paddle1+halfPaddle) {
			e.State.BallX = 3
			e.VelocityX = -e.VelocityX
		}
	}

	// check if ball in right
	if e.State.BallX >= 97 {
		if e.State.BallY >= (e.State.Paddle2-halfPaddle) && e.State.BallY <= (e.State.Paddle2+halfPaddle) {
			e.State.BallX = 97
			e.VelocityX = -e.VelocityX
		}
	}

	if e.State.BallX < 0 {
		e.resetBall()
	} else if e.State.BallX > 100 {
		e.resetBall()
	}

}

func (e *Engine) resetBall() {
	e.State.BallX = 50.0
	e.State.BallY = 50.0

	e.VelocityY = -e.VelocityY
}

func (e *Engine) MovePaddle(playerID int, direction string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	speed := 5.0

	if playerID == 1 {
		if direction == "UP" && e.State.Paddle1 > 10 {
			e.State.Paddle1 -= speed
		} else if direction == "DOWN" && e.State.Paddle1 < 90 {
			e.State.Paddle1 += speed
		}
	} else {
		if direction == "UP" && e.State.Paddle2 > 10 {
			e.State.Paddle2 -= speed
		} else if direction == "DOWN" && e.State.Paddle2 < 90 {
			e.State.Paddle2 += speed
		}
	}
}
