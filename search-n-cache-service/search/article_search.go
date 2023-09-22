package search

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSearch interface {
	GetArticleList() ([]*domain.ArticleListType, error)
	GetArticleQuery(string) ([]*domain.ArticleListType, error)
	SaveArticle(*domain.ArticleType) error
	UpdateArticle(*domain.ArticleType) error
	DeleteArticle(int64) error
}
