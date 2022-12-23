package user

import (
	"github.com/google/uuid"
)

// User is a aggregate root domain type
type User struct {
	UUID           uuid.UUID
	Email          Email
	HashedPassword HashedPassword
}

func New(id uuid.UUID, email Email, password HashedPassword) User {
	return User{
		UUID:           id,
		Email:          email,
		HashedPassword: password,
	}
}
