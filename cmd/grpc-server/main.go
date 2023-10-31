package main

import (
	"fmt"
	"log"
	"net"
	"os"

	api "github.com/Soumik43/grpc-user-service/api/user"
	user "github.com/Soumik43/grpc-user-service/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Get port from environment variables
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server started at port", port)

	userRepo := user.NewInMemoryUserRepository()
	grpcServer := grpc.NewServer()
	api.RegisterUserServiceServer(grpcServer, user.NewUserServiceServer(userRepo))
	
	// Register reflection service on gRPC server to use with grpcurl
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
