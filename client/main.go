package main

import (
	"context"
	"log"

	pb "test_ssl/proto/test_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverAddr := "localhost:50051" // Set the address of the gRPC server

	// Load SSL/TLS certificate
	certFile := "certs/localhost.pem" // Path to the PEM file containing the certificate and private key
	cred, err := credentials.NewClientTLSFromFile(certFile, "localhost")
	if err != nil {
		log.Fatalf("Failed to load SSL/TLS certificate: %v", err)
	}

	// Create a new gRPC connection with SSL/TLS
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewMyServiceClient(conn)

	// Send a gRPC request
	req := &pb.TestRequest{
		Message: "Hello, Server!",
	}

	response, err := client.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
