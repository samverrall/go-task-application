package grpc

import (
	"fmt"
	"net"

	"github.com/samverrall/go-task-application/logger"
	gen "github.com/samverrall/go-task-application/task-application-proto/gen"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"google.golang.org/grpc"
)

type GRPC struct {
	logger      logger.Logger
	userService user.API
	port        int
	server      *grpc.Server
	gen.UnimplementedUserServer
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
	g.server = grpcServer
	gen.RegisterUserServer(grpcServer, g)
	if err := grpcServer.Serve(listen); err != nil {
		g.logger.Error("failed to serve gRPC on port: %d, error: %v", g.port, err)
		return err
	}

	g.logger.Info("gRPC adapter listening on %d", g.port)

	return nil
}

func (g *GRPC) Stop() {
	g.server.Stop()
}
