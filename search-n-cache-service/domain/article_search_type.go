package domain

import (
	"time"
)

type ArticleSearchType struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Author          string    `json:"author"`
	ArticleCategory string    `json:"category"`
	ArticleType     string    `json:"type"`
	Tags            string    `json:"tags"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
}
