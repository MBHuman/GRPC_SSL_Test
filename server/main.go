package main

import (
	"context"
	"log"
	"net"
	pb "test_ssl/proto/test_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type MyServiceServer struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyServiceServer) MyMethod(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	// Implement the logic for your RPC method here
	response := &pb.TestResponse{
		Message: "Hello, " + req.Message,
	}
	return response, nil
}

func main() {
	listenAddr := "localhost:50051" // Set the address on which your server will listen

	// Load SSL/TLS certificate and key
	certFile := "certs/localhost.crt" // Replace with the path to your SSL/TLS certificate
	keyFile := "certs/localhost.key"  // Replace with the path to your private key
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load SSL/TLS credentials: %v", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMyServiceServer(grpcServer, &MyServiceServer{})

	log.Printf("Server listening on %s", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
