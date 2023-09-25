package search

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSolrSearch struct{}

func (s *ArticleSolrSearch) GetArticleList() ([]*domain.ArticleSearchType, error) {
	return nil, nil
}

func (s *ArticleSolrSearch) SearchArticles() ([]*domain.ArticleSearchType, error) {
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

func (s *ArticleSolrSearch) CreateIndex() error {
	return nil
}

func (s *ArticleSolrSearch) DeleteIndex() error {
	return nil
}
