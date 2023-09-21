package service

import (
	//"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/entity"
	"github.com/search-n-cache/search-n-cache-service/repository"
	log "github.com/sirupsen/logrus"
)

func GetArticleDetails(id string) *entity.Article{
	log.Info("Starting getArticleDetails")
	article := repository.GetArticle(id)
	return article
}
