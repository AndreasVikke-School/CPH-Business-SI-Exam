package main

import (
	"context"
	"encoding/json"
	"time"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	"github.com/segmentio/kafka-go"
)

func ProduceLogEntryToKafka(logEntry LogEntry) {
	topic := "logEntry"
	partition := 0

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := kafka.DialLeader(ctx, "tcp", configuration.Kafka.Service, topic, partition)
	eh.PanicOnError(err, "failed to dial leader")

	c, err := json.Marshal(logEntry)
	eh.PanicOnError(err, "Can't convert to JSON")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(c)})
	eh.PanicOnError(err, "failed to write messages")
}
