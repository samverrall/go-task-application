package user

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

type API interface {
	CreateUser(context.Context, user.User) error
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

func (us *UserService) CreateUser(ctx context.Context, u user.User) error {
	us.logger.Info("us.CreateUser Invoked")
	return us.repo.CreateUser(ctx, u)
}
