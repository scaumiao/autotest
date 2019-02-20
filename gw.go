package autotest

import (
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	taskProto "github.com/scaumiao/autotest/proto/task"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func newGateway(ctx context.Context, Endpoint string, opts ...gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := taskProto.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, Endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
