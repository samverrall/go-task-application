package grpc

import (
	"context"

	gen "github.com/samverrall/go-task-application/user-service/internal/adapters/left/grpc/gen/proto"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
)

func (g *GRPC) GetUserEmail(ctx context.Context, request *gen.GetUserEmailRequest) (*gen.GetUserEmailResponse, error) {
	g.logger.Info("GetUserEmail Invoked")

	user, err := g.userService.GetUser(ctx, user.GetUserRequest{
		UserUUID: request.UserUuid,
	})
	if err != nil {
		return nil, err
	}

	return &gen.GetUserEmailResponse{
		Email: user.Email,
	}, nil
}
