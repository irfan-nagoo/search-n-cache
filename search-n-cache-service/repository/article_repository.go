package repository

import (
	"github.com/search-n-cache/search-n-cache-service/entity"
	"gorm.io/gorm"
)

func CreateArticle(article *entity.Article, tx *gorm.DB) *gorm.DB {
	return tx.Create(article)
}

func GetArticle(id string, article *entity.Article, tx *gorm.DB) *gorm.DB {
	return tx.First(article, "id = ?", id)
}

func UpdateArticle(article *entity.Article, tx *gorm.DB) *gorm.DB {
	return tx.Updates(article)
}

func DeleteArticle(article *entity.Article, tx *gorm.DB) *gorm.DB {
	return tx.Delete(article)
}
