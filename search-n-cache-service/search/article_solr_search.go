package search

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSolrSearch string

func (s *ArticleSolrSearch) GetArticleList() ([]*domain.ArticleListType, error) {
	return nil, nil
}

func (s *ArticleSolrSearch) GetArticleQuery() ([]*domain.ArticleListType, error) {
	return nil, nil
}

func (s *ArticleSolrSearch) SaveArticle(articleType *domain.ArticleType) error {
	return nil
}

func (s *ArticleSolrSearch) UpdateArticle(articleType *domain.ArticleType) error {
	return nil
}

func (s *ArticleSolrSearch) DeleteArticle(id int64) error {
	return nil
}
