package main

import (
	"log"
	"net"
	"net/http"

	"github.com/sbasestarter/grpc-production-demo/backserver/internal/handlers"
	"github.com/sbasestarter/grpc-production-demo/backserver/internal/server"
	"github.com/sbasestarter/grpc-production-demo/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gRpcServer := grpc.NewServer()
	hellopb.RegisterHellosServer(gRpcServer, server.NewGrpcServer())
	reflection.Register(gRpcServer)

	go func() {
		gRpcListen, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("grpc server listen on:", gRpcListen.Addr())
		err = gRpcServer.Serve(gRpcListen)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	gRpcWebListen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	h, err := handlers.NewGRpcWebHandler(handlers.GRpcWebHandlerInputParameters{
		GrpcServer:          gRpcServer,
		GRpcWebUseWebsocket: false,
		GRpcWebPingInterval: 0,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("grpc web server listen on:", gRpcWebListen.Addr())
	httpServer := &http.Server{Handler: h}
	err = httpServer.Serve(gRpcWebListen)
	if err != nil {
		log.Fatalln(err)
	}
}
