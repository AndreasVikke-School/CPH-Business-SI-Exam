package main

import (
	"context"

	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/rpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetLog(ctx context.Context, in *pb.LogRequest) (*pb.Log, error) {
	log, err := GetLogFromRedis(in.UserId, in.Id, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: log.Id, UserId: log.UserId, EntityId: log.EntityId, Unix: log.Unix}, nil
}

func (s *server) GetAllLogs(ctx context.Context, in *wrapperspb.Int64Value) (*pb.LogList, error) {
	logs, err := GetAllLogsFromRedis(in.Value, configuration)
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
	createdLog, err := CreateLogInRedis(in, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: createdLog.Id, UserId: createdLog.UserId, EntityId: createdLog.EntityId, Unix: createdLog.Unix}, nil
}
