package domain

import (
	"time"
)


type ArticleType struct {
	ID int64
	Title string
	ArticleCategory string
	ArticleType string
	Content string
	Tags string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}