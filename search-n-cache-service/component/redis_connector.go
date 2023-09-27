package component

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func InitializeRedisCacheClient() *redis.Client {
	log.Info("Initializing Redis Client")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	log.Info("Redis Client Initialized Successfully")
	return redisClient
}
