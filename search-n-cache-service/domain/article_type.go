package domain

import (
	"github.com/search-n-cache/search-n-cache-service/constants"
	"time"
)

type ArticleType struct {
	ID              int64              `json:"id"`
	Title           string             `json:"title" binding:"required"`
	Description     string             `json:"description" binding:"required"`
	Author          string             `json:"author" binding:"required"`
	ArticleCategory constants.Category `json:"category"`
	ArticleType     constants.Type     `json:"type"`
	Content         string             `json:"content" binding:"required"`
	Tags            string             `json:"tags"`
	CreatedAt       time.Time          `json:"createdAt"`
	CreatedBy       string             `json:"createdBy"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	UpdatedBy       string             `json:"updatedBy"`
}
