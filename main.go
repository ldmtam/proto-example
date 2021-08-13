package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ldmtam/proto-example/internal/greeting"
	greetingv1 "github.com/ldmtam/proto-example/proto/greeting/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	greetingv1.RegisterGreeterServer(grpcServer, greeting.NewServer())
	reflection.Register(grpcServer)

	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = greetingv1.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:3000", opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8081", mux)
}
