package repository

import (
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/entity"
)

func GetArticle(id string) *entity.Article {
	article := &entity.Article{}
	component.DBConnection.First(article, "id = ?", id)
	return article
}
