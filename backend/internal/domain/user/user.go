package user

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             UserID
	Email          Email
	HashedPassword Password
	Name           Name
	CreatedAt      time.Time
}

type UserID string
type Email string
type Password string
type Name string

func NewUser(email Email, rawPassword string, name Name, hasher PasswordHasher) (*User, error) {
	if err := email.Validate(); err != nil {
		return nil, err
	}

	if err := name.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(rawPassword, hasher)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             UserID(uuid.NewString()),
		Email:          email,
		HashedPassword: Password(hashedPassword),
		Name:           name,
		CreatedAt:      time.Now().UTC(),
	}, nil
}

func (u *User) ComparePassword(rawPassword string, hasher PasswordHasher) bool {
	return hasher.Compare(string(u.HashedPassword), rawPassword)
}

func (u *User) UpdatePassword(newPassword string, hasher PasswordHasher) error {
	hashed, err := hashPassword(newPassword, hasher)
	if err != nil {
		return err
	}
	u.HashedPassword = Password(hashed)
	return nil
}

func (e Email) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(string(e)) {
		return ErrInvalidEmail
	}
	return nil
}

func (n Name) Validate() error {
	if strings.TrimSpace(string(n)) == "" {
		return ErrInvalidName
	}
	return nil
}

func hashPassword(rawPassword string, hasher PasswordHasher) (string, error) {
	if len(rawPassword) < 8 {
		return "", ErrInvalidPassword
	}

	return hasher.Hash(rawPassword)
}
