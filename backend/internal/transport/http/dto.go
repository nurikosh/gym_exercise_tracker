package http

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,gte=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SessionResponse struct {
	ID        int                `json:"id"`
	Date      string             `json:"date"`
	Notes     string             `json:"notes"`
	Exercises []ExerciseResponse `json:"exercises"`
}

type ExerciseResponse struct {
	Name string        `json:"name"`
	Sets []SetResponse `json:"sets"`
}

type SetResponse struct {
	Weight float64 `json:"weight"`
	Reps   int     `json:"reps"`
}
