package gateway

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/task-application-proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	logger logger.Logger
}

// New returns a pointer to a new Gateway
func New(log logger.Logger) *Gateway {
	return &Gateway{
		logger: log,
	}
}

func (g *Gateway) Handler(ctx context.Context, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// TODO: Remove hard coded endpoints, read from a config instead.
	err := gen.RegisterUserHandlerFromEndpoint(ctx, mux, "127.0.0.1:8000", dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
