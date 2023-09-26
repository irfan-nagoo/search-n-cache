package domain

import (
	"github.com/search-n-cache/search-n-cache-service/constants"
	"time"
)

type ArticleSearchType struct {
	ID              int64              `json:"id"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Author          string             `json:"author"`
	ArticleCategory constants.Category `json:"category"`
	ArticleType     constants.Type     `json:"type"`
	Tags            string             `json:"tags"`
	CreatedAt       time.Time          `json:"createdAt"`
	CreatedBy       string             `json:"createdBy"`
}
