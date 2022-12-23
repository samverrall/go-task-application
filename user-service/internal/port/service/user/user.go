package user

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/pkg/hasher"
)

type API interface {
	GetUser(context.Context, GetUserRequest, GetUserGuard) (*getUserResponse, error)
	Register(context.Context, RegisterRequest) error
}

type UserService struct {
	repo   domain.UserRepo
	logger logger.Logger
	hasher hasher.Hasher
}

func NewService(repo domain.UserRepo, logger logger.Logger, hasher hasher.Hasher) API {
	return &UserService{
		repo:   repo,
		logger: logger,
		hasher: hasher,
	}
}
