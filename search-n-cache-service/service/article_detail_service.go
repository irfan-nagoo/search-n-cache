package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/cache"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/constants"
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/entity"
	"github.com/search-n-cache/search-n-cache-service/error"
	"github.com/search-n-cache/search-n-cache-service/repository"
	"github.com/search-n-cache/search-n-cache-service/request"
	"github.com/search-n-cache/search-n-cache-service/response"
	"github.com/search-n-cache/search-n-cache-service/search"
	log "github.com/sirupsen/logrus"
)

// Save Article
//
//	@Summary      Save article
//	@Description  Save Article Information
//	@Tags         Article Detail
//	@Accept       json
//	@Produce      json
//	@Param 		  request body  request.ArticleRequest true "Article request"
//	@Success      200  {object}  response.ArticleDetailResponse
//	@Failure      400  {object}  response.ErrorResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article [post]
func SaveArticle(ctx *gin.Context) {
	log.Info("Starting SaveArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		error.HandleError(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.CreateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	articleType := component.ArticleToArticleType(article)
	// Create or Index the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).SaveArticle(articleType); err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(articleType, ctx)
}

// Get Article
//
//	@Summary      Get article
//	@Description  Get Article Information
//	@Tags         Article Detail
//	@Accept       json
//	@Produce      json
//	@Param 		  id   path      int  true  "Article Id"
//	@Success      200  {object}  response.ArticleDetailResponse
//	@Failure      404  {object}  response.ErrorResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article/{id} [get]
func GetArticle(ctx *gin.Context) {
	log.Info("Starting GetArticleDetails")
	id := ctx.Params.ByName("id")

	// try with cache first
	articleCache := cache.ArticleRedisCache{}
	cachedArticle, _ := articleCache.ReadFromCache("article:" + id)
	if cachedArticle != nil {
		buildArtDetlSuccessResponse(cachedArticle, ctx)
		return
	}

	// if not, read from database and set in cache
	article := &entity.Article{}
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.GetArticle(id, article, tx).Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	// write to cache
	articleType := component.ArticleToArticleType(article)
	if err := articleCache.WriteToCache("article:"+id, articleType); err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(articleType, ctx)
}

// Update Article
//
//	@Summary      Update article
//	@Description  Update Article Information
//	@Tags         Article Detail
//	@Accept       json
//	@Produce      json
//	@Param 		  request body  request.ArticleRequest true "Article request"
//	@Success      200  {object}  response.ArticleDetailResponse
//	@Failure      400  {object}  response.ErrorResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article [put]
func UpdateArticle(ctx *gin.Context) {
	log.Info("Starting UpdateArticle")
	var articleRequest *request.ArticleRequest
	if err := ctx.BindJSON(&articleRequest); err != nil {
		error.HandleError(err, ctx)
		return
	}
	article := component.ArticleTypeToArticle(articleRequest.Article)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.UpdateArticle(article, tx).Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	// Update the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).UpdateArticle(articleRequest.Article); err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(component.ArticleToArticleType(article), ctx)
}

// Delete Article
//
//	@Summary      Delete article
//	@Description  Delete Article Information
//	@Tags         Article Detail
//	@Accept       json
//	@Produce      json
//	@Param 		  id   path      int  true  "Article Id"
//	@Success      200  {object}  response.BaseResponse
//	@Failure      404  {object}  response.ErrorResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article/{id} [delete]
func DeleteArticle(ctx *gin.Context) {
	log.Info("Starting DeleteArticle")
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		error.HandleError(err, ctx)
	}
	articleType := &domain.ArticleType{ID: id}
	article := component.ArticleTypeToArticle(articleType)
	// start new transaction
	tx := component.DBConnection.Begin()
	if err := repository.DeleteArticle(article, tx).Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	// Delete the article record in the Search Engine
	if err := (&search.ArticleElasticSearch{}).DeleteArticle(id); err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		error.HandleError(err, ctx)
		return
	}
	// transaction completed
	buildArtDetlSuccessResponse(nil, ctx)
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
