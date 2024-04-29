package main

import (
	"context"
	"time"

	// codes "google.golang.org/grpc/codes"
	// status "google.golang.org/grpc/status"

	"task_service/proto"
)

type taskServiceServer struct {
	proto.UnimplementedTaskServiceServer

	taskDB DBHandler
}

func GetProtoTask(task Task) *proto.Task {
	return &proto.Task{TaskId: task.TaskId, Name: task.Name, Description: task.Description, DeadlineDate: task.DeadlineDate, CreationDate: task.CreationDate, CompletionStatus: task.CompletionStatus}
}

func (server *taskServiceServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	taskID, err := server.taskDB.CreateTask(Task{AuthorId: req.UserId, Name: req.Name, Description: req.Description,
		DeadlineDate: req.DeadlineDate, CreationDate: time.Now().Format("2017.09.07 17:06:06")})
	return &proto.CreateTaskResponse{TaskId: taskID}, err
}

func (server *taskServiceServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.UpdateTaskResponse, error) {
	err := server.taskDB.UpdateTask(Task{AuthorId: req.UserId, TaskId: req.TaskId, Name: req.Name, Description: req.Description,
		DeadlineDate: req.DeadlineDate})
	return &proto.UpdateTaskResponse{}, err
}

func (server *taskServiceServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	err := server.taskDB.DeleteTask(req.UserId, req.TaskId)
	return &proto.DeleteTaskResponse{}, err
}

func (server *taskServiceServer) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	task, err := server.taskDB.GetTask(req.UserId, req.TaskId)
	return &proto.GetTaskResponse{Task: GetProtoTask(task)}, err
}

func (server *taskServiceServer) GetTasksPage(ctx context.Context, req *proto.GetTasksPageRequest) (*proto.GetTasksPageResponse, error) {
	tasks, err := server.taskDB.GetTasksPage(req.UserId, req.PageIndex, req.TasksPerPage)
	protoTasks := make([]*proto.Task, 0, len(tasks))
	for _, task := range tasks {
		protoTasks = append(protoTasks, GetProtoTask(task))
	}
	return &proto.GetTasksPageResponse{Tasks: protoTasks}, err
}
