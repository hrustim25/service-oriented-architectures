package main

import (
	"context"

	"stat_service/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type statServiceServer struct {
	proto.UnimplementedStatServiceServer

	statDB DBHandler
}

func (server *statServiceServer) GetEventsCount(ctx context.Context, in *proto.GetEventsCountRequest) (*proto.GetEventsCountResponse, error) {
	viewCount, err := server.statDB.GetEventCountForTask(in.TaskId, ViewEventID)
	if err != nil {
		return nil, err
	}
	likeCount, err := server.statDB.GetEventCountForTask(in.TaskId, LikeEventID)
	if err != nil {
		return nil, err
	}
	return &proto.GetEventsCountResponse{ViewCount: viewCount, LikeCount: likeCount}, nil
}

func (server *statServiceServer) GetTopTasks(ctx context.Context, in *proto.GetTopTasksRequest) (*proto.GetTopTasksResponse, error) {
	topTasks, err := server.statDB.GetTopTasks(in.EventType)
	if err != nil {
		return nil, err
	}
	resp := &proto.GetTopTasksResponse{}
	resp.Tasks = make([]*proto.TopTask, 0, len(topTasks))
	for _, task := range topTasks {
		resp.Tasks = append(resp.Tasks, &proto.TopTask{TaskId: task.TaskId, TaskAuthorId: task.TaskAuthorId, ViewCount: task.ViewCount, LikeCount: task.LikeCount})
	}
	return resp, nil
}

func (server *statServiceServer) GetTopAuthors(ctx context.Context, in *emptypb.Empty) (*proto.GetTopAuthorsResponse, error) {
	topAuthors, err := server.statDB.GetTopAuthors()
	if err != nil {
		return nil, err
	}
	resp := &proto.GetTopAuthorsResponse{}
	resp.Authors = make([]*proto.TopAuthor, 0, len(topAuthors))
	for _, author := range topAuthors {
		resp.Authors = append(resp.Authors, &proto.TopAuthor{TaskAuthorId: author.AuthorId, LikeCount: author.LikeCount})
	}
	return resp, nil
}
