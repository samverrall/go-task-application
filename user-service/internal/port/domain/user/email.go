package user

import (
	"errors"
	"strings"
)

type Email string

func NewEmail(email string) (Email, error) {
	if !strings.Contains(email, "@") {
		return "", errors.New("email should cotain an @ symbol")
	}
	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}
