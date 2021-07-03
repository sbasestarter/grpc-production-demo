package handlers

import (
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRpcWebHandler struct {
	gRpcServer          *grpc.Server
	gRpcWebUseWebsocket bool
	gRpcWebPingInterval time.Duration
}

type GRpcWebHandlerInputParameters struct {
	GrpcServer          *grpc.Server
	GRpcWebUseWebsocket bool
	GRpcWebPingInterval time.Duration
}

func NewGRpcWebHandler(parameters GRpcWebHandlerInputParameters) (http.Handler, error) {
	if parameters.GrpcServer == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	return &gRpcWebHandler{
		gRpcServer:          parameters.GrpcServer,
		gRpcWebUseWebsocket: parameters.GRpcWebUseWebsocket,
		gRpcWebPingInterval: parameters.GRpcWebPingInterval,
	}, nil
}

func (s *gRpcWebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	}
	if s.gRpcWebUseWebsocket {
		options = append(
			options,
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		)

		if s.gRpcWebPingInterval > 0 {
			options = append(options, grpcweb.WithWebsocketPingInterval(s.gRpcWebPingInterval))
		}
	}
	wrappedGrpc := grpcweb.WrapServer(s.gRpcServer, options...)
	wrappedGrpc.ServeHTTP(w, r)
}
