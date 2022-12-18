package main

import (
	"context"
	"fmt"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/task-application-proto/gen"
)

func newGateway(ctx context.Context, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	err := gen.RegisterUserHandlerFromEndpoint(ctx)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log := logger.New("debug")

	gw, err := newGateway(ctx, nil)
	if err != nil {
		log.Fatal("failed to init new gateway: %s", err.Error())
	}

	fmt.Println(gw)

	// TODO: Init new HTTP server
	// TODO: Attach newGateway handler to mux
}
