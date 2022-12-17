package user

import (
	"github.com/google/uuid"
)

// User is a aggregate root domain type
type User struct {
	UUID     uuid.UUID
	Email    Email
	Password Password
}

func New(id uuid.UUID, email Email, password Password) User {
	return User{
		UUID:     id,
		Email:    email,
		Password: password,
	}
}
