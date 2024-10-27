package main

import (
	pb "grpc-project/proto"
	services "grpc-project/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	log.Println("Hello, World!")
	grpcServer := grpc.NewServer()
	testService := services.NewTestService()
	pb.RegisterHelloServiceServer(grpcServer, testService)

	// Lắng nghe trên cổng 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
