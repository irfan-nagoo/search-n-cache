package response

import (
	"github.com/search-n-cache/search-n-cache-service/domain"
)

type ArticleSearchResponse struct {
	BaseResponse
	Articles []*domain.ArticleSearchType `json:"articles,omitempty"`
}
