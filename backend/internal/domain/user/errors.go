package user

import "errors"

var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 8 characters")
	ErrInvalidName     = errors.New("name cannot be empty")
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailTaken      = errors.New("email already registered")
)
