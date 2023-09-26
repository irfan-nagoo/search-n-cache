package component

import (
	"crypto/tls"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func InitializeRedisCacheClient() *redis.Client {
	log.Info("Initializing Redis Client")
	skipVerifySSL, _ := strconv.ParseBool(os.Getenv("REDIS_SKIP_VERIFY_SSL"))
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: skipVerifySSL,
		},
	})
	log.Info("Redis Client Initialzed Successfully")
	return redisClient
}
