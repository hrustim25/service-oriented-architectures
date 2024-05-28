package main

import (
	"log"
	"net"
	"net/http"

	"stat_service/proto"

	"google.golang.org/grpc"
)

func main() {
	SetupDB()
	SetupHandlers()

	SetupAndStartStatMessageBrokerConsumer()

	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		proto.RegisterStatServiceServer(grpcServer, &statServiceServer{statDB: statDB})

		log.Default().Println("Starting gRPC stat server...")

		grpcServer.Serve(listener)
	}()

	log.Default().Println("Starting stat server...")

	err := http.ListenAndServe("0.0.0.0:12345", nil)
	if err != nil {
		panic("Server falled")
	}
}
