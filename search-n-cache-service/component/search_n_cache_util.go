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

func ArticleTypeToArticleSearchType(articleType *domain.ArticleType) *domain.ArticleSearchType {
	articleSearchType := domain.ArticleSearchType{
		ID:              articleType.ID,
		Title:           articleType.Title,
		Description:     articleType.Description,
		Author:          articleType.Author,
		ArticleCategory: articleType.ArticleCategory,
		ArticleType:     articleType.ArticleType,
		Tags:            articleType.Tags,
		CreatedAt:       articleType.CreatedAt,
		CreatedBy:       articleType.CreatedBy,
	}
	return &articleSearchType
}
