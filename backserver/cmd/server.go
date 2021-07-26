package main

import (
	"log"
	"net"
	"net/http"

	"github.com/sbasestarter/grpc-production-demo/backserver/internal/handlers"
	"github.com/sbasestarter/grpc-production-demo/backserver/internal/server"
	"github.com/sbasestarter/grpc-production-demo/proto/gen/go"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	go func() {
		// /debug/rpcz /debug/tracez
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe("127.0.0.1:8089", mux))
	}()

	view.RegisterExporter(&exporter.PrintExporter{})
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	gRpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	hellopb.RegisterHellosServer(gRpcServer, server.NewGrpcServer())
	reflection.Register(gRpcServer)

	go func() {
		gRpcListen, err := net.Listen("tcp", ":9080")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("grpc server listen on:", gRpcListen.Addr())
		err = gRpcServer.Serve(gRpcListen)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	gRpcWebListen, err := net.Listen("tcp", ":8082")
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
