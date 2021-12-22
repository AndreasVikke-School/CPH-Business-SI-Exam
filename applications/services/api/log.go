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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Log struct {
	Id       int64 `json:"id,omitempty"`
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	Unix     int64 `json:"unix,omitempty"`
}

type LogEntry struct {
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	UnixTime int64 `json:"unix,omitempty"`
}

// Get Log By User
// @Schemes
// @Description  Gets a log by user
// @Tags         Log
// @Accept       json
// @Param        userId  path  int  true  "Id of user"
// @Param        logId  path  int  true  "Id of log"
// @Produce      json
// @Success      200  {object}  Log
// @Failure      404
// @Router       /log/get-by-user/{userId}/{logId} [get]
func GetLogByUser(c *gin.Context) {
	userId := c.Param("userId")
	uId, err := strconv.ParseInt(userId, 10, 64)
	eh.PanicOnError(err, "failed to parse userId to int64")

	logId := c.Param("logId")
	lId, err := strconv.ParseInt(logId, 10, 64)
	eh.PanicOnError(err, "failed to parse logId to int64")

	conn, err := grpc.Dial(configuration.Redis.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log, err := con.GetLogByUser(ctx, &pb.LogRequest{Id: lId, UserId: uId})
	if err != nil {
		eh.PanicOnError(err, "error")
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, log)
	}
}

// Get All Logs By User
// @Schemes
// @Description  Says a list of all logs by user
// @Tags         Log
// @Accept       json
// @Param        id  path  int  true  "Id of user"
// @Produce      json
// @Success      200  {object}  []Log
// @Failure      404
// @Router       /log/all-by-user/{id} [get]
func GetAllLogsByUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	eh.PanicOnError(err, "failed to parse userId to int64")

	conn, err := grpc.Dial(configuration.Redis.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logs, err := con.GetAllLogsByUser(ctx, &wrapperspb.Int64Value{Value: id})
	logList := []Log{}
	for _, l := range logs.Logs {
		logList = append(logList, Log{Id: l.Id, UserId: l.UserId, EntityId: l.EntityId, Unix: l.Unix})
	}

	if err != nil {
		eh.PanicOnError(err, "error")
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, logList)
	}
}

// Create Log Entry
// @Schemes
// @Description  Creates a log entry
// @Tags         Log
// @Accept       json
// @Param        LogEntry  body  LogEntry  true  "Create log entry"
// @Produce      json
// @Success      200
// @Router       /log/create/ [post]
func CreateLogEntry(c *gin.Context) {
	var logEntry LogEntry
	err := c.BindJSON(&logEntry)
	eh.PanicOnError(err, "Couldn't bind JSON")
	logEntry.UnixTime = time.Now().UnixNano() / 1000000
	ProduceLogEntryToKafka(logEntry)
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
