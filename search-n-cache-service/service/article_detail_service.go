package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/constants"
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/entity"
	"github.com/search-n-cache/search-n-cache-service/repository"
	"github.com/search-n-cache/search-n-cache-service/request"
	"github.com/search-n-cache/search-n-cache-service/response"
	"github.com/search-n-cache/search-n-cache-service/search"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SaveArticle(ctx *gin.Context) {
	log.Info("Starting SaveArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.CreateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	articleType := component.ArticleToArticleType(article)
	// Create or Index the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).SaveArticle(articleType); err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(articleType, ctx)
}

func GetArticle(ctx *gin.Context) {
	log.Info("Starting GetArticleDetails")
	id := ctx.Params.ByName("id")
	article := &entity.Article{}
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.GetArticle(id, article, tx).Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(component.ArticleToArticleType(article), ctx)
}

func UpdateArticle(ctx *gin.Context) {
	log.Info("Starting UpdateArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.BindJSON(&articleRequest); err != nil {
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.UpdateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	// Update the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).UpdateArticle(articleRequest.Article); err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(component.ArticleToArticleType(article), ctx)
}

func DeleteArticle(ctx *gin.Context) {
	log.Info("Starting DeleteArticle")
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		buildArtDetlErrorResponse(err, ctx)
	}
	articleType := &domain.ArticleType{ID: id}
	article := component.ArticleTypeToArticle(articleType)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.DeleteArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	// Delete the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).DeleteArticle(id); err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildArtDetlErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(nil, ctx)
}

func buildArtDetlErrorResponse(err error, ctx *gin.Context) {
	log.Error(err)
	httpStatus := http.StatusInternalServerError
	// specific error check
	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpStatus = http.StatusNotFound
	}
	// error type check
	if errors.As(err, &validator.ValidationErrors{}) {
		httpStatus = http.StatusBadRequest
	}
	res := response.BaseResponse{
		Code:    http.StatusText(httpStatus),
		Message: err.Error(),
	}
	ctx.JSON(httpStatus, res)
}

func buildArtDetlSuccessResponse(articleType *domain.ArticleType, ctx *gin.Context) {
	httpStatus := http.StatusOK
	res := response.ArticleDetailResponse{
		BaseResponse: response.BaseResponse{
			Code:    http.StatusText(httpStatus),
			Message: constants.SUCCESS_MESSAGE},
		Article: articleType,
	}
	ctx.JSON(httpStatus, res)
}
