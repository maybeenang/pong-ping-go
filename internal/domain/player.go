// Package domain
package domain

import "time"

type Player struct {
	ID        string
	Username  string
	CreatedAt time.Time
}
