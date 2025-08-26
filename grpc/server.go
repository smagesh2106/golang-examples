package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc/hellopb"

	"google.golang.org/grpc"
)

type server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	msg := fmt.Sprintf("Hello, %s, Age:%s !", req.GetName(), req.GetAge())
	return &hellopb.HelloResponse{Message: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(grpcServer, &server{})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
