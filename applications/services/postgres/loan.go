package main

import (
	"context"

	pb "github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetLoan(ctx context.Context, in *wrapperspb.Int64Value) (*pb.Loan, error) {
	l, u, err := PostgresGetLoanById(context.Background(), in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.Loan{Id: int64(l.ID), UserId: int64(u.ID), EntityId: l.EntityId, Status: pb.Status(pb.Status_value[string(l.Status)])}, nil
}

func (s *server) GetAllLoans(ctx context.Context, in *emptypb.Empty) (*pb.LoanList, error) {
	loans, users, err := PostgresGetAllLoans(context.Background(), configuration)
	if err != nil {
		return nil, err
	}

	var result []*pb.Loan
	for idx, l := range loans {
		result = append(result, &pb.Loan{Id: int64(l.ID), UserId: int64(users[idx].ID), EntityId: l.EntityId, Status: pb.Status(pb.Status_value[string(l.Status)])})
	}
	return &pb.LoanList{Loans: result}, nil
}

func (s *server) CreateLoan(ctx context.Context, in *pb.Loan) (*pb.Loan, error) {
	createdLoan, err := PostgresInsertNewLoan(ctx, in, configuration)

	if err != nil {
		return nil, err
	}
	return &pb.Loan{Id: int64(createdLoan.ID), UserId: in.UserId, EntityId: createdLoan.EntityId, Status: in.Status}, nil
}
