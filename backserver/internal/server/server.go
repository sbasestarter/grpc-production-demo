package server

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/sbasestarter/grpc-production-demo/proto/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRpcServer struct {
	idx int64
}

func (s *gRpcServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	ret := fmt.Sprintf("%s-%d", req.GetRequest(), atomic.AddInt64(&s.idx, 1))
	return &hellopb.HelloResponse{
		Response: ret,
	}, nil
}

func NewGrpcServer() hellopb.HellosServer {
	return &gRpcServer{}
}
