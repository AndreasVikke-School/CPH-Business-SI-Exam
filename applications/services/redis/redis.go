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
	Id       int64 `json:"id,omitempty"`
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

	id, err := rdb.HLen(rdb.Context(), redis_key).Result()
	if err != nil {
		log.Printf("Hlen error: %s", err)
		return 0, nil, err
	}

	unix := time.Now().UnixNano() / 1000000
	dataAsJson := fmt.Sprintf(`{"userId": %d, "entryId": %d, "unix": %d}`, in.UserId, in.EntityId, unix)

	_, err = rdb.HSet(rdb.Context(), redis_key, (id + 1), dataAsJson).Result()
	if err != nil {
		log.Printf("Hset error: %s", err)
		return 0, nil, err
	}

	return id + 1, &LogEntry{UserId: in.UserId, EntityId: in.EntityId, Unix: unix}, nil
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
		log.Printf("HGet error: %s", err)
		return 0, nil, err
	}

	var data *LogEntry
	json.Unmarshal([]byte(result), &data)

	return logId, data, nil
}

func GetAllLogsFromRedis(config Configuration) ([]*LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	result, err := rdb.HGetAll(rdb.Context(), redis_key).Result()
	if err != nil {
		log.Printf("HGet error: %s", err)
		return nil, err
	}
	fmt.Println(result)

	var logs []*LogEntry
	for _, res := range result {
		var data *LogEntry
		json.Unmarshal([]byte(res), &data)
		logs = append(logs, data)
	}

	return logs, nil
}
