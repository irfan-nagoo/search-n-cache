package cache

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/domain"
	log "github.com/sirupsen/logrus"
)

type ArticleRedisCache struct{}

func (c *ArticleRedisCache) ReadFromCache(key string) (*domain.ArticleType, error) {
	article := domain.ArticleType{}
	if len(key) > 0 {
		// get record
		record, err := component.RedisClient.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}

		log.Infof("Cache Hit for key[%s]", key)
		if err := json.Unmarshal([]byte(record), &article); err != nil {
			return nil, err
		}
	}
	return &article, nil
}

func (c *ArticleRedisCache) WriteToCache(key string, article *domain.ArticleType) error {
	if len(key) > 0 && article != nil {
		log.Infof("Set Cache Record for key[%s]", key)
		cacheExpirySec, _ := strconv.ParseInt(os.Getenv("REDIS_CACHE_EXPIRY_INTERVAL_SEC"), 10, 64)
		articleBytes, err := json.Marshal(article)
		if err != nil {
			return err
		}
		// set with expiry
		component.RedisClient.Set(context.Background(), key, articleBytes, time.Duration(cacheExpirySec)*time.Second)
	}
	return nil
}
