package storage

import "time"

type UserModel struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type SessionModel struct {
	ID     int       `db:"id"`
	UserID int       `db:"user_id"`
	Date   time.Time `db:"date"`
	Notes  string    `db:"notes"`
}

type ExerciseModel struct {
	ID        int    `db:"id"`
	SessionID int    `db:"session_id"`
	Name      string `db:"name"`
}

type SetModel struct {
	ID         int     `db:"id"`
	ExerciseID int     `db:"exercise_id"`
	Weight     float32 `db:"weight"`
	Reps       int     `db:"reps"`
	SetOrder   int     `db:"set_order"`
}
