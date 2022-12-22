package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/samverrall/go-task-application/logger"
	gen "github.com/samverrall/go-task-application/task-application-proto/gen"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPC struct {
	logger      logger.Logger
	userService user.API
	port        int
	host        string
	server      *grpc.Server
	gen.UnimplementedUserServer
}

func New(userSvc user.API, logger logger.Logger, host string, port int) *GRPC {
	return &GRPC{
		logger:      logger,
		userService: userSvc,
		port:        port,
		host:        host,
	}
}

func (g *GRPC) Run(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", g.host, g.port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		g.logger.Error("failed to listen on port %d, error: %v", g.port, err)
		return err
	}
	defer func() {
		if err := listen.Close(); err != nil {
			g.logger.Error("Failed to close %s %s: %v", "tcp", addr, err)
		}
	}()

	g.server = grpc.NewServer()
	gen.RegisterUserServer(g.server, g)
	reflection.Register(g.server)

	go func() {
		defer g.server.GracefulStop()
		<-ctx.Done()
	}()

	g.logger.Info("gRPC adapter listening on %s", addr)
	if err := g.server.Serve(listen); err != nil {
		g.logger.Error("failed to serve gRPC on port: %d, error: %v", g.port, err)
		return err
	}

	return nil
}

func (g *GRPC) Stop() {
	g.server.Stop()
}
