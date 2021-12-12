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
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	Unix     int64 `json:"unix,omitempty"`
}

func GetRedisClient(config Configuration) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{config.Redis.Broker},
		Password: "",
	})
}

func CreateLogInRedis(in *pb.Log, config Configuration) (*LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	userKey := fmt.Sprintf(`%s:%d`, redis_key, in.UserId)

	id, err := rdb.HLen(rdb.Context(), userKey).Result()
	if err != nil {
		log.Printf("Hlen error: %s", err)
		return nil, err
	}

	unix := time.Now().UnixNano() / 1000000
	dataAsJson := fmt.Sprintf(`{"entityId": %d, "unix": %d}`, in.EntityId, unix)

	_, err = rdb.HSet(rdb.Context(), userKey, (id + 1), dataAsJson).Result()
	if err != nil {
		log.Printf("Hset error: %s", err)
		return nil, err
	}

	return &LogEntry{Id: id + 1, UserId: in.UserId, EntityId: in.EntityId, Unix: unix}, nil
}

func GetLogFromRedis(userId int64, logId int64, config Configuration) (*LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	userKey := fmt.Sprintf(`%s:%d`, redis_key, userId)

	exists := rdb.HExists(rdb.Context(), userKey, strconv.FormatInt(logId, 10)).Val()
	if !exists {
		log.Printf("log not found in redis")
		return nil, errors.New("log not found in redis")
	}

	result, err := rdb.HGet(rdb.Context(), userKey, strconv.FormatInt(logId, 10)).Result()
	if err != nil {
		log.Printf("HGet error: %s", err)
		return nil, err
	}

	var data *LogEntry
	json.Unmarshal([]byte(result), &data)
	data.UserId = userId
	data.Id = logId

	return data, nil
}

func GetAllLogsFromRedis(userId int64, config Configuration) ([]*LogEntry, error) {
	rdb := GetRedisClient(config)
	defer rdb.Close()

	userKey := fmt.Sprintf(`%s:%d`, redis_key, userId)

	result, err := rdb.HGetAll(rdb.Context(), userKey).Result()
	if err != nil {
		log.Printf("HGet error: %s", err)
		return nil, err
	}
	fmt.Println(result)

	var logs []*LogEntry
	for key, res := range result {
		var data *LogEntry
		json.Unmarshal([]byte(res), &data)
		id, _ := strconv.ParseInt(key, 10, 64)
		data.UserId = userId
		data.Id = id
		logs = append(logs, data)
	}

	return logs, nil
}
