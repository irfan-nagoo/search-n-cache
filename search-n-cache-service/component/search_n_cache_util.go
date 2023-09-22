package component

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/entity"
)

func ArticleTypeToArticle(articleType *domain.ArticleType) *entity.Article {
	article := entity.Article(*articleType)
	return &article
}

func ArticleToArticleType(article *entity.Article) *domain.ArticleType {
	articleType := domain.ArticleType(*article)
	return &articleType
}
