package user

import (
	"github.com/google/uuid"
)

const (
	PasswordMinLength = 6
)

// User is a aggregate root domain type
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
