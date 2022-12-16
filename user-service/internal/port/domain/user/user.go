package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	PasswordMinLength = 6
)

type User struct {
	ID       uuid.UUID
	Email    Email
	Password Password
}

func New(id uuid.UUID, email Email, password Password) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}

}

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

type Password string

func NewPassword(password string) (Password, error) {
	if len(password) < PasswordMinLength {
		return "", fmt.Errorf("a password must be greater than %d characters", PasswordMinLength)
	}
	return Password(password), nil
}
