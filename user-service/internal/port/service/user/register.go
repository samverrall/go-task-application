package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

type RegisterRequest struct {
	Email    string
	Password string
}

func (us *UserService) Register(ctx context.Context, in RegisterRequest) error {
	log.Info("userService.Register Invoked")

	password, err := user.NewPassword(in.Password)
	if err != nil {
		return err
	}

	email, err := user.NewEmail(in.Email)
	if err != nil {
		return err
	}

	uuid := uuid.New()

	newUser := user.New(uuid, email, password)

	return us.repo.Add(ctx, newUser)
}
