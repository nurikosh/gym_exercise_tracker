package user

import (
	"errors"
	"gym/internal/domain"
	"gym/internal/domain/ports"
	"time"
)

type User struct {
	ID           domain.UserID
	Email        Email
	PasswordHash string
	Name         string
	CreatedAt    time.Time
}

func NewUser(email Email, passwordHash, name string, hasher ports.PasswordHasher) (*User, error) {
	if !email.IsValid() {
		return nil, ErrInvalidEmail
	}

	if name == "" {
		return nil, ErrNameRequired
	}

	passwordHash, err := hasher.Hash(passwordHash)
	if err != nil {
		return nil, errors.New("Error hashing password")
	}

	return &User{
		ID:           domain.NewID(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		CreatedAt:    time.Now(),
	}, nil
}

func (u *User) UpdateEmail(newEmail Email) error {
	if !newEmail.IsValid() {
		return ErrInvalidEmail
	}
	u.Email = newEmail
	return nil
}
func (u *User) UpdateName(newName string) error {
	if newName == "" {
		return ErrNameRequired
	}
	u.Name = newName
	return nil
}

func (u *User) UpdatePassword(newPassword string, hasher ports.PasswordHasher) error {
	if newPassword == "" {
		return ErrWeakPassword
	}

	passwordHash, err := hasher.Hash(newPassword)
	if err != nil {
		return errors.New("Error hashing password")
	}
	u.PasswordHash = passwordHash
	return nil
}

func (u *User) CheckPassword(password string, hasher ports.PasswordHasher) bool {
	return hasher.Compare(u.PasswordHash, password)
}
