package gateway

import (
	"context"
	"fmt"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samverrall/go-task-application/gateway-service/config"
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/task-application-proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	logger logger.Logger
	config *config.Config
}

// New returns a pointer to a new Gateway
func New(log logger.Logger, c *config.Config) *Gateway {
	return &Gateway{
		logger: log,
		config: c,
	}
}

func (g *Gateway) Handler(ctx context.Context, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userServiceAddr := buildAddress(g.config.GetString("user-service.host"), g.config.GetInt("user-service.port"))
	err := gen.RegisterUserHandlerFromEndpoint(ctx, mux, userServiceAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func buildAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
