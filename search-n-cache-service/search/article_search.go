package search

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSearch interface {
	
	GetArticleList(int,int) ([]*domain.ArticleSearchType, error)
	SearchArticles(string,int,int) ([]*domain.ArticleSearchType, error)
	SaveArticle(*domain.ArticleType) error
	UpdateArticle(*domain.ArticleType) error
	DeleteArticle(int64) error
	CreateIndex() error
	DeleteIndex() error
}
