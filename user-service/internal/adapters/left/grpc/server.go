package grpc

import (
	"fmt"
	"net"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"google.golang.org/grpc"
)

type GRPC struct {
	logger      logger.Logger
	userService user.API
	port        int
}

func New(userSvc user.API, logger logger.Logger, port int) *GRPC {
	return &GRPC{
		logger:      logger,
		userService: userSvc,
		port:        port,
	}
}

func (g *GRPC) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		g.logger.Error("failed to listen on port %d, error: %v", g.port, err)
		return err
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listen); err != nil {
		g.logger.Error("failed to serve gRPC on port: %d, error: %v", g.port, err)
		return err
	}

	return nil
}
