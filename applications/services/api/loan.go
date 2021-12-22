package main

import (
	"context"
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

type LoanEntry struct {
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
}

// Get Loan
// @Schemes
// @Description  Gets a loan by id
// @Tags         Loan
// @Accept       json
// @Param        id  path  int  true  "Id of loan"
// @Produce      json
// @Success      200  {object}  Loan
// @Failure      404
// @Router       /loan/get/{id} [get]
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
		eh.PanicOnError(err, "error")
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, loan)
	}
}

// Get All Loans
// @Schemes
// @Description  Gets a list of all loans
// @Tags         Loan
// @Accept       json
// @Produce      json
// @Success      200  {object}  []Loan
// @Failure      404
// @Router       /loan/all/ [get]
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
		eh.PanicOnError(err, "error")
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, loansList)
	}
}

// Get All Loans By User
// @Schemes
// @Description  Gets a list of all loans by a user
// @Tags         Loan
// @Accept       json
// @Param        id  path  int  true  "Id of all loans by user"
// @Produce      json
// @Success      200  {object}  []Loan
// @Failure      404
// @Router       /loan/all-by-user/{id} [get]
func GetAllLoansByUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	eh.PanicOnError(err, "failed to parse loanId to int64")

	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLoanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	loans, err := con.GetAllLoansByUser(ctx, &wrapperspb.Int64Value{Value: id})
	loansList := []Loan{}
	for _, l := range loans.Loans {
		loansList = append(loansList, Loan{Id: l.Id, UserId: l.UserId, EntityId: l.EntityId, Status: l.Status})
	}

	if err != nil {
		eh.PanicOnError(err, "error")
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, loansList)
	}
}

// Create Loan Entry
// @Schemes
// @Description  Creates a loan entry
// @Tags         Loan
// @Accept       json
// @Param        LoanEntry  body  LoanEntry  true  "Create loan"
// @Produce      json
// @Success      200
// @Router       /loan/create/ [post]
func CreateLoanEntry(c *gin.Context) {
	var loanEntry LoanEntry
	err := c.BindJSON(&loanEntry)
	eh.PanicOnError(err, "Couldn't bind JSON")
	ProduceLoanEntryToRabbitmq(loanEntry)
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
