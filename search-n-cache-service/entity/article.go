package entity

import (
	"time"
)


type Article struct {
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

type Tabler interface {
	TableName() string
}

func (Article) TableName() string {
	return "sc_article"
} 