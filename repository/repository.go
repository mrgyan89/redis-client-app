package repository

import (
	"context"
	"fmt"
	"github.com/mrgyan89/redis-client-app/model"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client
var redisConnectionString string = "redis://user:pass@<host>:6379"
var ctx = context.Background()

// Non TLS connection
func init() {
	opt, err := redis.ParseURL(redisConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	redisClient = redis.NewClient(opt)
	fmt.Println("Redis connection established successfully")
}

func GetMessages() []model.UserEntry {
	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	var userEntries []model.UserEntry
	for _, key := range keys {
		fmt.Println(key)
		userEntry := GetMessageByKey(key)
		userEntries = append(userEntries, userEntry)
	}
	return userEntries
}
func GetMessageByKey(key string) model.UserEntry {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return model.UserEntry{Key: key, Value: val}
}

func SetMessage(userEntry model.UserEntry) model.UserEntry {
	err := redisClient.Set(ctx, userEntry.Key, userEntry.Value, 0).Err()
	if err != nil {
		panic(err)
	}
	return userEntry
}

func DeleteByKey(key string) {
	redisClient.Del(ctx, key)
}
