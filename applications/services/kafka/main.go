package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	eh "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/kafka/errorhandler"
	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/kafka/rpc"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

var configuration Configuration

type LogEntry struct {
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	Unix     int64 `json:"unix,omitempty"`
}

func consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{configuration.Kafka.Broker},
		Topic:          "logEntry",
		GroupID:        "kafka",
		MinBytes:       1,    // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
	})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
		os.Exit(1)
	}()

	fmt.Println("Starting reading from queue")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		var data *LogEntry
		json.Unmarshal([]byte(m.Value), &data)
		CreateLogEntryInRedis(data)
	}
}

func CreateLogEntryInRedis(data *LogEntry) {
	conn, err := grpc.Dial(configuration.Redis.Broker, grpc.WithInsecure())
	eh.PanicOnError(err, "failed connecting to grpc")
	defer conn.Close()

	c := pb.NewLogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.CreateLog(ctx, &pb.Log{UserId: data.UserId, EntityId: data.EntityId, Unix: data.Unix})
	eh.PanicOnError(err, "failed to create log")
	log.Printf("Log created in redis with id: %d", r.Id)
}

func main() {

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	consume()
}
