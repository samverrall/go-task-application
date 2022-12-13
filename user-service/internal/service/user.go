package service

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/domain"
)

type UserServicer interface {
	CreateUser(context.Context, *domain.User) error
}

type UserService struct {
	repo   domain.UserRepo
	logger logger.Logger
}

func NewUserService(repo domain.UserRepo, logger logger.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (us *UserService) CreateUser(ctx context.Context, u *domain.User) error {
	us.logger.Info("us.CreateUser Invoked")
	return us.repo.CreateUser(ctx, u)
}
