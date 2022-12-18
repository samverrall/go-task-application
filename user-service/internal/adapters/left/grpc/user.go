package grpc

import (
	"context"

	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
)

type canUserGetEmail struct{}

// TODO: Implement proper check
func (u *canUserGetEmail) CanGetUser() bool {
	return true
}

func (g *GRPC) GetUserEmail(ctx context.Context, request *gen.GetUserEmailRequest) (*gen.GetUserEmailResponse, error) {
	g.logger.Info("GetUserEmail Invoked")

	user, err := g.userService.GetUser(ctx, user.GetUserRequest{
		UserUUID: request.UserUuid,
	}, &canUserGetEmail{})
	if err != nil {
		g.logger.Error("failed to get user: %s", err.Error())
		return nil, err
	}

	return &gen.GetUserEmailResponse{
		Email: user.Email,
	}, nil
}
