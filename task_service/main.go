package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"task_service/proto"
)

const (
	DB_URL_ENV = "POSTGRES_URL"
)

func main() {
	taskDB := SetupDB()

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaskServiceServer(grpcServer, &taskServiceServer{taskDB: taskDB})

	log.Default().Println("Starting gRPC task server...")

	grpcServer.Serve(listener)
}
