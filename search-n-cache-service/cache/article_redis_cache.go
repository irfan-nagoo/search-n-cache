package cache

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleRedisCache struct{}

var ctx = context.Background()

func (c *ArticleRedisCache) ReadFromCache(key string) (*domain.ArticleType, error) {
	article := domain.ArticleType{}
	if len(key) > 0 {
		if articleCmd := component.RedisClient.Get(ctx, key); articleCmd != nil {
			articleBytes, _ := articleCmd.Bytes()
			if err := json.Unmarshal(articleBytes, &article); err != nil {
				return nil, err
			}
		}
	}
	return &article, nil
}
func (c *ArticleRedisCache) WriteToCache(key string, article *domain.ArticleType) error {
	if len(key) > 0 && article != nil {
		cacheExpirySec, _ := strconv.ParseInt(os.Getenv("REDIS_CACHE_EXPIRY_INTERVAL_SEC"), 10, 16)
		articleBytes, err := json.Marshal(article)
		if err != nil {
			return err
		}
		component.RedisClient.Set(ctx, key, articleBytes, time.Duration(cacheExpirySec))
	}
	return nil
}
