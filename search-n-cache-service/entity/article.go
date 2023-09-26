package entity

import (
	"time"
	"github.com/search-n-cache/search-n-cache-service/constants"
)

type Article struct {
	ID              int64
	Title           string
	Description     string
	Author          string
	ArticleCategory constants.Category
	ArticleType     constants.Type
	Content         string
	Tags            string
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
}

type Tabler interface {
	TableName() string
}

func (Article) TableName() string {
	return "sc_article"
}
