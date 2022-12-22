package grpc

import (
	"context"
	"errors"

	gen "github.com/samverrall/go-task-application/task-application-proto/gen"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type canUserGetEmail struct{}

// TODO: Implement proper check
func (u *canUserGetEmail) CanGetUser() bool {
	return true
}

func (g *GRPC) GetUserEmail(ctx context.Context, request *gen.GetUserEmailRequest) (*gen.GetUserEmailResponse, error) {
	g.logger.Info("GRPC.GetUserEmail Invoked")

	user, err := g.userService.GetUser(ctx, user.GetUserRequest{
		UserUUID: request.UserUuid,
	}, &canUserGetEmail{})
	switch {
	case errors.Is(err, repository.ErrUserNotFound):
		err := grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "404"))
		if err != nil {
			g.logger.LogError(err)
		}
		g.logger.Info("Set gRPC http header")

		return nil, err

	case err != nil:
		g.logger.Error("failed to get user: %s", err.Error())
		return nil, err
	}

	return &gen.GetUserEmailResponse{
		Email: user.Email,
	}, nil
}
