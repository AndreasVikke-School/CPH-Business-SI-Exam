package main

import (
	"context"

	pb "github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/rpc"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetUser(ctx context.Context, in *wrapperspb.Int64Value) (*pb.User, error) {
	u, err := PostgresGetUserById(context.Background(), in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.User{Id: int64(u.ID), Username: u.Username, Age: u.Age}, nil
}

func (s *server) GetAllUsers(ctx context.Context, in *emptypb.Empty) (*pb.UserList, error) {
	users, err := PostgresGetAllUsers(context.Background(), configuration)
	if err != nil {
		return nil, err
	}

	var result []*pb.User
	for _, u := range users {
		result = append(result, &pb.User{Id: int64(u.ID), Username: u.Username, Age: u.Age, Password: u.Password})
	}
	return &pb.UserList{Users: result}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	createdUser, err := PostgresInsertNewUser(ctx, in, configuration)

	if err != nil {
		return nil, err
	}
	return &pb.User{Id: int64(createdUser.ID), Username: createdUser.Username, Age: createdUser.Age, Password: createdUser.Password}, nil
}
