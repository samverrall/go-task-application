package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

type API interface {
	GetUser(context.Context, GetUserDTO) (*user.User, error)
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

type GetUserDTO struct {
	UserUUID string
}

func (us *UserService) GetUser(ctx context.Context, userDTO GetUserDTO) (*user.User, error) {
	us.logger.Info("us.CreateUser Invoked")

	uuid, err := uuid.Parse(userDTO.UserUUID)
	if err != nil {
		return nil, err
	}

	return us.repo.Get(ctx, uuid)
}
