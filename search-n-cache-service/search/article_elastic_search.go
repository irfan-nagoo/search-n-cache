package search

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/domain"
	log "github.com/sirupsen/logrus"
)

type ArticleElasticSearch struct{}

func (s *ArticleElasticSearch) GetArticleList() ([]*domain.ArticleSearchType, error) {
	log.Info("Starting GetArticleList")
	searchArticles := []*domain.ArticleSearchType{}
	searchResult, err := component.ElasticTypedClient.
		Search().
		Index(os.Getenv("ELASTIC_INDEX")).
		Request(
			&search.Request{
				Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
			}).
		Do(context.TODO())

	if err != nil {
		log.Errorf("Error occured while invoking Elastic Search API: %s", err)
		return nil, err
	}
	log.Infof("Time taken: %d ms", searchResult.Took)
	if searchResult.Hits.Total.Value > 0 {

		for _, hit := range searchResult.Hits.Hits {
			article := domain.ArticleSearchType{}
			err := json.Unmarshal(hit.Source_, &article)
			if err != nil {
				log.Errorf("Error occured while deserialization: %s", err)
				return nil, err
			}
			searchArticles = append(searchArticles, &article)
		}

	}
	return searchArticles, nil
}

func (s *ArticleElasticSearch) SearchArticles() ([]*domain.ArticleSearchType, error) {
	return nil, nil
}

func (s *ArticleElasticSearch) SaveArticle(articleType *domain.ArticleType) error {
	log.Info("Starting SaveArticle")
	article := component.ArticleTypeToArticleSearchType(articleType)
	_, err := component.ElasticTypedClient.
		Index(os.Getenv("ELASTIC_INDEX")).
		Id(strconv.FormatInt(article.ID, 16)).
		Request(article).
		Do(context.TODO())

	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleElasticSearch) UpdateArticle(articleType *domain.ArticleType) error {
	log.Info("Starting UpdateArticle")
	article := component.ArticleTypeToArticleSearchType(articleType)
	json, er := json.Marshal(article)
	if er != nil {
		log.Errorf("Error occured while serialization: %s", er)
		return er
	}

	_, err := component.ElasticTypedClient.
		Update(os.Getenv("ELASTIC_INDEX"), strconv.FormatInt(article.ID, 16)).
		Request(&update.Request{
			Doc: json,
		}).
		Do(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleElasticSearch) DeleteArticle(id int64) error {
	log.Info("Starting DeleteArticle")
	_, err := component.ElasticTypedClient.
		Delete(os.Getenv("ELASTIC_INDEX"), strconv.FormatInt(id, 16)).
		Do(context.TODO())

	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleElasticSearch) CreateIndex() error {
	log.Info("Starting CreateIndex")
	res, err := component.ElasticTypedClient.
		Indices.
		Exists(os.Getenv("ELASTIC_INDEX")).
		Do(context.TODO())
	if !res {
		log.Info("Creating Index")
		component.ElasticTypedClient.
			Indices.
			Create(os.Getenv("ELASTIC_INDEX")).
			Do(context.TODO())
	}
	return err
}

func (s *ArticleElasticSearch) DeleteIndex() error {
	log.Info("Starting DeleteIndex")
	res, err := component.ElasticTypedClient.Indices.
		Exists(os.Getenv("ELASTIC_INDEX")).
		Do(context.TODO())
	if res {
		log.Info("Deleting Index")
		component.ElasticTypedClient.Indices.
			Delete(os.Getenv("ELASTIC_INDEX")).
			Do(context.TODO())
	}
	return err
}
