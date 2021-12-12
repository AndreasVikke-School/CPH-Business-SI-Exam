package main

import (
	"context"

	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetLog(ctx context.Context, in *wrapperspb.Int64Value) (*pb.Log, error) {
	id, log, err := GetLogFromRedis(in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: id, UserId: log.UserId, EntityId: log.EntityId, Unix: log.Unix}, nil
}

func (s *server) GetAllLoans(ctx context.Context, in *emptypb.Empty) (*pb.LogList, error) {
	logs, err := GetAllLogsFromRedis(configuration)
	if err != nil {
		return nil, err
	}

	var result []*pb.Log
	for _, l := range logs {
		result = append(result, &pb.Log{Id: l.Id, UserId: l.UserId, EntityId: l.EntityId, Unix: l.Unix})
	}
	return &pb.LogList{Logs: result}, nil
}

func (s *server) CreateLog(ctx context.Context, in *pb.Log) (*pb.Log, error) {
	id, createdLog, err := CreateLogInRedis(in, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: id, UserId: createdLog.UserId, EntityId: createdLog.EntityId, Unix: createdLog.Unix}, nil
}
