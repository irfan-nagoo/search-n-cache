package cache

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleMemCache struct{}

func (c *ArticleMemCache) ReadFromCache(key string) (*domain.ArticleType, error) {
	return nil, nil
}
func (c *ArticleMemCache) WriteToCache(key string, article *domain.ArticleType) error {
	return nil
}
