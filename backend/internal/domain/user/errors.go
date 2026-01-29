package user

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidEmail   = errors.New("invalid email format")
	ErrWeakPassword   = errors.New("password does not meet requirements")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrNameRequired   = errors.New("name is required")
)
