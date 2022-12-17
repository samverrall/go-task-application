package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type GetUserGuard interface {
	CanGetUser() bool
}

// GetUserRequest is a DTO that adapters have to pass to the port.
// This helps to keep all domain validation on the core, and non-adapter
// specific.
type GetUserRequest struct {
	UserUUID string
}

type GetUserResponse struct {
	ID       string
	Email    string
	Password string
}

func (us *UserService) GetUser(ctx context.Context, userDTO GetUserRequest, guard GetUserGuard) (*GetUserResponse, error) {
	us.logger.Info("us.CreateUser Invoked")

	if authorised := guard.CanGetUser(); !authorised {
		return nil, errors.New("not allowed to get user")
	}

	uuid, err := uuid.Parse(userDTO.UserUUID)
	if err != nil {
		return nil, err
	}

	got, err := us.repo.Get(ctx, uuid)
	if err != nil {

	}

	return &GetUserResponse{
		ID:    got.ID.String(),
		Email: got.Email.String(),
	}, nil
}
