package domain

import "time"

type Session struct {
	ID        int64
	UserID    int64
	Date      time.Time
	Notes     string
	Exercises []Exercise
}
