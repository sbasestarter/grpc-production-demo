package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/sbasestarter/grpc-production-demo/proto/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRpcServer struct {
	idx      int64
	hostName string
}

func (s *gRpcServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	ret := fmt.Sprintf("%s:%s-%d", s.hostName, req.GetRequest(), atomic.AddInt64(&s.idx, 1))
	return &hellopb.HelloResponse{
		Response: ret,
	}, nil
}

func (s *gRpcServer) HelloStream(req *hellopb.HelloStreamRequest, stream hellopb.Hellos_HelloStreamServer) error {
	log.Println("HelloStream enter")
	if req == nil || !strings.HasPrefix(req.Auth, "123") {
		log.Println("HelloStream leave: no auth")
		return status.Error(codes.Unauthenticated, "need auth,兄弟")
	}
	var err error
	var idx int
	for {
		err = stream.Send(&hellopb.HelloStreamMessage{
			Message: fmt.Sprintf("%s:%s-%d", s.hostName, req.Auth, idx),
		})
		if err != nil {
			break
		}
		idx++
		time.Sleep(time.Second)
	}
	log.Println("HelloStream leave:", err)
	return err
}

func NewGrpcServer() hellopb.HellosServer {
	host, err := os.Hostname()
	if err != nil {
		host = err.Error()
	}
	return &gRpcServer{
		hostName: host,
	}
}
