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

type User struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Age      int64  `json:"age,omitempty"`
	Password string `json:"password,omitempty"`
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	eh.PanicOnError(err, "failed to parse userId to int64")

	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := con.GetUser(ctx, &wrapperspb.Int64Value{Value: id})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func GetAllUsers(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	users, err := con.GetAllUsers(ctx, &emptypb.Empty{})
	usersList := []User{}
	for _, u := range users.Users {
		usersList = append(usersList, User{Id: u.Id, Username: u.Username, Age: u.Age})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, usersList)
	}
}

func CreateUser(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var user User
	err = c.BindJSON(&user)
	eh.PanicOnError(err, "Couldn't bind json to user")

	usr, err := con.CreateUser(ctx, &pb.User{Id: user.Id, Username: user.Username, Age: user.Age, Password: user.Password})
	eh.PanicOnError(err, "failed to create user")
	log.Printf("User created in postgres with id: %d", usr.Id)

}
