package domain

import "time"

type Session struct {
	ID        int
	UserID    int
	Date      time.Time
	Notes     string
	Exercises []Exercise
}
