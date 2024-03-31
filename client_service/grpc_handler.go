package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client_service/proto"
)

type TaskServiceHandler struct {
	client proto.TaskServiceClient
}

var taskService TaskServiceHandler

func SetupTaskServiceHandler() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("task_service:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to grpc dial: %v", err)
	}
	taskService = TaskServiceHandler{client: proto.NewTaskServiceClient(conn)}
}
