package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/constants"
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/entity"
	"github.com/search-n-cache/search-n-cache-service/repository"
	"github.com/search-n-cache/search-n-cache-service/request"
	"github.com/search-n-cache/search-n-cache-service/response"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SaveArticle(ctx *gin.Context) {
	log.Info("Starting SaveArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.BindJSON(&articleRequest); err != nil {
		buildErrorResponse(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.CreateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildSuccessResponse(component.ArticleToArticleType(article), ctx)
}

func GetArticle(ctx *gin.Context) {
	log.Info("Starting GetArticleDetails")
	id := ctx.Params.ByName("id")
	article := &entity.Article{}
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.GetArticle(id, article, tx).Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildSuccessResponse(component.ArticleToArticleType(article), ctx)
}

func UpdateArticle(ctx *gin.Context) {
	log.Info("Starting UpdateArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.BindJSON(&articleRequest); err != nil {
		buildErrorResponse(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.UpdateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildSuccessResponse(component.ArticleToArticleType(article), ctx)
}

func DeleteArticle(ctx *gin.Context) {
	log.Info("Starting DeleteArticle")
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		buildErrorResponse(err, ctx)
	}
	articleType := &domain.ArticleType{ID: id}
	article := component.ArticleTypeToArticle(articleType)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.DeleteArticle(article, tx).Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		buildErrorResponse(err, ctx)
		return
	}
	// transaction completed
	buildSuccessResponse(nil, ctx)
}

func buildErrorResponse(err error, ctx *gin.Context) {
	log.Error(err)
	httpStatus := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpStatus = http.StatusNotFound
	}
	res := response.BaseResponse{
		Code:    http.StatusText(httpStatus),
		Message: err.Error(),
	}
	ctx.JSON(httpStatus, res)
}

func buildSuccessResponse(articleType *domain.ArticleType, ctx *gin.Context) {
	httpStatus := http.StatusOK
	res := response.ArticleResponse{
		BaseResponse: response.BaseResponse{
			Code:    http.StatusText(httpStatus),
			Message: constants.SUCCESS_MESSAGE},
		Article: articleType,
	}
	ctx.JSON(httpStatus, res)
}
