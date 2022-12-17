package user

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
)

type API interface {
	GetUser(context.Context, GetUserRequest, GetUserGuard) (*getUserResponse, error)
	Register(context.Context, RegisterRequest) error
}

type UserService struct {
	repo   domain.UserRepo
	logger logger.Logger
}

func NewService(repo domain.UserRepo, logger logger.Logger) API {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}
