package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent"
	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent/loan"
	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent/user"
	pb "github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/rpc"

	_ "github.com/lib/pq"
)

func GetPostgresClient(config Configuration) *ent.Client {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=disable", config.Postgres.Broker, config.Postgres.Port, config.Postgres.User, config.Postgres.Password)
	fmt.Println(conn)
	client, err := ent.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func MigratePostgres(config Configuration) {
	client := GetPostgresClient(config)
	defer client.Close()

	err := client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// ==== User Postgres ====
func PostgresInsertNewUser(ctx context.Context, in *pb.User, config Configuration) (*ent.User, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	u, err := client.User.
		Create().
		SetUsername(in.Username).
		SetAge(in.Age).
		SetPassword(in.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)
	return u, nil
}

func PostgresGetUserById(ctx context.Context, id int64, config Configuration) (*ent.User, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	u, err := client.User.
		Query().
		Where(user.ID(int(id))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)
	return u, nil
}

func PostgresGetAllUsers(ctx context.Context, config Configuration) ([]*ent.User, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	u, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}

	log.Println("users returned: ", u)
	return u, nil
}

// ==== User Postgres END ====

// ==== Loan Postgres ====
func PostgresInsertNewLoan(ctx context.Context, in *pb.Loan, config Configuration) (*ent.Loan, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	u, err := client.User.
		Query().
		Where(user.ID(int(in.UserId))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	l, err := client.Loan.
		Create().
		SetEntityId(in.EntityId).
		SetStatus(loan.Status(in.Status.String())).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating loan: %w", err)
	}

	u.Update().AddLoans(l).Save(ctx)

	log.Println("loan was created: ", l)
	return l, nil
}

func PostgresGetLoanById(ctx context.Context, id int64, config Configuration) (*ent.Loan, *ent.User, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	l, err := client.Loan.
		Query().
		Where(loan.ID(int(id))).
		Only(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed querying loan: %w", err)
	}

	u, err := l.QueryUser().Only(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("loan returned: ", l)
	return l, u, nil
}

func PostgresGetAllLoans(ctx context.Context, config Configuration) ([]*ent.Loan, []*ent.User, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	loans, err := client.Loan.Query().All(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed querying users: %w", err)
	}

	var users []*ent.User
	for _, l := range loans {
		users = append(users, l.QueryUser().OnlyX(ctx))
	}

	log.Println("loans returned: ", loans)
	return loans, users, nil
}

// ==== Loan Postgres END ====
