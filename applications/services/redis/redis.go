package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/rpc"
	"github.com/go-redis/redis/v8"
)

var redis_key = "log_entry"

type LogEntry struct {
	UserId   int64 `json:"studentId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	Unix     int64 `json:"unix,omitempty"`
}

func GetRedisClient(config Configuration) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{config.Redis.Broker},
		Password: "",
	})
}

func CreateLogInRedis(in *pb.Log, config Configuration) (int64, *LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	id, err := rdb.Incr(rdb.Context(), redis_key).Result()
	if err != nil {
		log.Printf("%s", err)
		return 0, nil, err
	}

	unix := time.Now().UnixNano() / 1000000
	dataAsJson := fmt.Sprintf(`{"userId": %d, "entryId": %d, "unix": %d}`, in.UserId, in.EntityId, unix)

	_, err = rdb.HSet(rdb.Context(), redis_key, strconv.FormatInt(id, 10), dataAsJson).Result()
	if err != nil {
		log.Printf("%s", err)
		return 0, nil, err
	}

	return id, &LogEntry{UserId: in.UserId, EntityId: in.EntityId, Unix: unix}, nil
}

func GetLogFromRedis(logId int64, config Configuration) (int64, *LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	exists := rdb.HExists(rdb.Context(), redis_key, strconv.FormatInt(logId, 10)).Val()
	if !exists {
		log.Printf("log not found in redis")
		return 0, nil, errors.New("log not found in redis")
	}

	result, err := rdb.HGet(rdb.Context(), redis_key, strconv.FormatInt(logId, 10)).Result()
	if err != nil {
		log.Printf("%s", err)
		return 0, nil, err
	}

	var data *LogEntry
	json.Unmarshal([]byte(result), &data)

	return logId, data, nil
}