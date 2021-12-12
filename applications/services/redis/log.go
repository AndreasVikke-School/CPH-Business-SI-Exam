package main

import (
	"context"

	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/rpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetLog(ctx context.Context, in *wrapperspb.Int64Value) (*pb.Log, error) {
	id, log, err := GetLogFromRedis(in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: id, UserId: log.UserId, EntityId: log.EntityId, Unix: log.Unix}, nil
}

// func (s *server) GetAllLoans(ctx context.Context, in *emptypb.Empty) (*pb.LoanList, error) {
// 	loans, users, err := PostgresGetAllLoans(context.Background(), configuration)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result []*pb.Loan
// 	for idx, l := range loans {
// 		result = append(result, &pb.Loan{Id: int64(l.ID), UserId: int64(users[idx].ID), EntityId: l.EntityId, Status: pb.Status(pb.Status_value[string(l.Status)])})
// 	}
// 	return &pb.LoanList{Loans: result}, nil
// }

func (s *server) CreateLog(ctx context.Context, in *pb.Log) (*pb.Log, error) {
	id, createdLog, err := CreateLogInRedis(in, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: id, UserId: createdLog.UserId, EntityId: createdLog.EntityId, Unix: createdLog.UserId}, nil
}
