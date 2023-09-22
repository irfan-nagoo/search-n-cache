package search

import (
	"context"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/search-n-cache/search-n-cache-service/domain"
	log "github.com/sirupsen/logrus"
)

type ArticleElasticSearch string

func (s *ArticleElasticSearch) GetArticleList() ([]*domain.ArticleListType, error) {
	log.Info("Starting GetArticleList")
	elasticTypedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_URL")}})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	elasticTypedClient.Search().Index(os.Getenv("ELASTIC_INDEX")).Request(&search.Request{
		Query: &types.Query{MatchAll: &types.MatchAllQuery{}}}).DO(context.TODO())
	return nil, nil
}

func (s *ArticleElasticSearch) GetArticleQuery() ([]*domain.ArticleListType, error) {
	return nil, nil
}

func (s *ArticleElasticSearch) SaveArticle(articleType *domain.ArticleType) error {
	return nil
}

func (s *ArticleElasticSearch) UpdateArticle(articleType *domain.ArticleType) error {
	return nil
}

func (s *ArticleElasticSearch) DeleteArticle(id int64) error {
	return nil
}
