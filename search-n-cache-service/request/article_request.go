package request

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleRequest struct {
	Article *domain.ArticleType `json:"article"`
}

func (artRequest ArticleRequest) SetArticle(article *domain.ArticleType) {
	artRequest.Article = article
}
