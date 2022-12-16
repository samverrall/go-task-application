package domain

import (
	"context"

	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

type UserRepo interface {
	CreateUser(context.Context, user.User) error
}
