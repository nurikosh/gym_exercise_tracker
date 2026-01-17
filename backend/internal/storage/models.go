package storage

import "time"

type UserModel struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type SessionModel struct {
	ID     int64     `db:"id"`
	UserID int64     `db:"user_id"`
	Date   time.Time `db:"date"`
	Notes  string    `db:"notes"`
}

type ExerciseModel struct {
	ID        int64  `db:"id"`
	SessionID int64  `db:"session_id"`
	Name      string `db:"name"`
}

type SetModel struct {
	ID         int64   `db:"id"`
	ExerciseID int64   `db:"exercise_id"`
	Weight     float64 `db:"weight"`
	Reps       int     `db:"reps"`
	SetOrder   int     `db:"set_order"`
}
