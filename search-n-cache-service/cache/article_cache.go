package cache


import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleCache interface {

	ReadFromCache(string) (*domain.ArticleType, error)
	WriteToCache(string, *domain.ArticleType) error
}