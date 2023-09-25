package search

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSearch interface {
	GetArticleList() ([]*domain.ArticleSearchType, error)
	SearchArticles(string) ([]*domain.ArticleSearchType, error)
	SaveArticle(*domain.ArticleType) error
	UpdateArticle(*domain.ArticleType) error
	DeleteArticle(int64) error
	CreateIndex() error
	DeleteIndex() error
}
