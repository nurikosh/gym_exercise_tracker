package domain

import (
	"time"
)

// User represents an authenticated gym member
type User struct {
	ID        int64
	Email     string
	Password  string // hashed
	Name      string
	CreatedAt time.Time
}
