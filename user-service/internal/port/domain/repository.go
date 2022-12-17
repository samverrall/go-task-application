package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

type UserRepo interface {
	Get(context.Context, uuid.UUID) (*user.User, error)
	Add(context.Context, user.User) error
}
