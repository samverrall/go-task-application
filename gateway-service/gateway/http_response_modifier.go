package gateway

import (
	"context"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
)

func (g *Gateway) httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	g.logger.Info("g.httpResponseModifier Invoked")

	md, ok := gwruntime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	spew.Dump(md)

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			g.logger.LogError(err)
			return err
		}
		g.logger.Debug("Status Code From gRPC: %d", code)

		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}
