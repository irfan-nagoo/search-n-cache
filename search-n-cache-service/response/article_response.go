package response

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleResponse struct {
	BaseResponse
	Article *domain.ArticleType `json:"article"`
}
