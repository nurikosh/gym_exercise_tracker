package user

import (
	"regexp"
)

type Email string

func (e Email) IsValid() bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(string(e))
}
