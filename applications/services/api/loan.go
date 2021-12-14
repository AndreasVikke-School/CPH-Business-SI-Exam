package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	pb "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Loan struct {
	Id       int64     `json:"id,omitempty"`
	UserId   int64     `json:"userId,omitempty"`
	EntityId int64     `json:"entityId,omitempty"`
	Status   pb.Status `json:"status,omitempty"`
}

func GetLoan(c *gin.Context) {
	loanId := c.Param("id")
	id, err := strconv.ParseInt(loanId, 10, 64)
	eh.PanicOnError(err, "failed to parse loanId to int64")

	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLoanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	loan, err := con.GetLoan(ctx, &wrapperspb.Int64Value{Value: id})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, loan)
	}
}

func GetAllLoans(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLoanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	loans, err := con.GetAllLoans(ctx, &emptypb.Empty{})
	loansList := []Loan{}
	for _, l := range loans.Loans {
		loansList = append(loansList, Loan{Id: l.Id, UserId: l.UserId, EntityId: l.EntityId, Status: l.Status})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, loansList)
	}
}

// func GetAllLoansByUser(c *gin.Context) {
// 	userId := c.Param("id")
// 	id, err := strconv.ParseInt(userId, 10, 64)
// 	eh.PanicOnError(err, "failed to parse loanId to int64")

// 	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
// 	eh.PanicOnError(err, "failed to connect to grpc")
// 	defer conn.Close()

// 	con := pb.NewLoanServiceClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	loans, err := con.GetAllLoansByUser(ctx, &emptypb.Empty{})
// 	// loansList := []Loan{}
// 	println(id)
// 	for _, l := range loans.Loans {
// 		println(l.UserId)
// 		// if l.UserId == id {
// 		// 	loansList = append(loansList, Loan{Id: l.Id, UserId: l.UserId, EntityId: l.EntityId, Status: l.Status})
// 		// }
// 	}

// 	// if err != nil {
// 	// 	c.Status(http.StatusNotFound)
// 	// } else {
// 	// 	c.IndentedJSON(http.StatusOK, loansList)
// 	// }
// }

func CreateLoan(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLoanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var loan Loan
	err = c.BindJSON(&loan)
	eh.PanicOnError(err, "Couldn't bind json to loan")

	ln, err := con.CreateLoan(ctx, &pb.Loan{Id: loan.Id, UserId: loan.UserId, EntityId: loan.EntityId, Status: loan.Status})
	eh.PanicOnError(err, "failed to create loan")
	log.Printf("Loan created in postgres with id: %d", ln.Id)
}
