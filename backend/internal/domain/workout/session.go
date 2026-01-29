package workout

import (
	"gym/internal/domain"
	"time"
)

type Session struct {
	ID     int
	UserID domain.UserID
	Date   time.Time
	Notes  string
}
