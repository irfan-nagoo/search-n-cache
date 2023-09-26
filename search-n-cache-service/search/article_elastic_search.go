package search

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/domain"
	log "github.com/sirupsen/logrus"
)

type ArticleElasticSearch struct{}

func (s *ArticleElasticSearch) GetArticleList(pageNo int, pageSize int) ([]*domain.ArticleSearchType, error) {
	log.Info("Starting GetArticleList")
	searchResult, err := component.ElasticTypedClient.
		Search().
		Index(os.Getenv("ELASTIC_INDEX")).
		Request(
			&search.Request{
				Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
				Sort: []types.SortCombinations{
					types.SortOptions{
						SortOptions: map[string]types.FieldSort{
							"createdAt": {Order: &sortorder.Desc},
						},
					},
				},
			}).
		From(pageNo).
		Size(pageSize).
		Do(context.TODO())
	return processSearchResult(searchResult, err)
}

func (s *ArticleElasticSearch) SearchArticles(query string, pageNo int, pageSize int) ([]*domain.ArticleSearchType, error) {
	log.Info("Starting SearchArticles")
	searchResult, err := component.ElasticTypedClient.
		Search().
		Index(os.Getenv("ELASTIC_INDEX")).
		Request(
			&search.Request{
				Query: &types.Query{
					MultiMatch: &types.MultiMatchQuery{
						Query:  query,
						Fields: []string{"title", "description", "author", "category", "type", "tags"},
					},
				},
			}).
		From(pageNo).
		Size(pageSize).
		Do(context.TODO())
	return processSearchResult(searchResult, err)
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

func processSearchResult(searchResult *search.Response, err error) ([]*domain.ArticleSearchType, error) {
	searchArticles := []*domain.ArticleSearchType{}
	if err != nil {
		return nil, err
	}
	log.Infof("Time taken: %d ms", searchResult.Took)
	if searchResult.Hits.Total.Value > 0 {

		for _, hit := range searchResult.Hits.Hits {
			article := domain.ArticleSearchType{}
			if err := json.Unmarshal(hit.Source_, &article); err != nil {
				return nil, err
			}
			searchArticles = append(searchArticles, &article)
		}

	}
	return searchArticles, nil
}
