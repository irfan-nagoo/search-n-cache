package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/constants"
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/response"
	"github.com/search-n-cache/search-n-cache-service/search"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetArticleList(ctx *gin.Context) {
	search := search.ArticleElasticSearch{}
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	response, err := search.GetArticleList(pageNo, pageSize)
	if err != nil {
		buildArtSrchErrorResponse(err, ctx)
		return
	}
	buildArtSrchSuccessResponse(response, ctx)
}

func SearchArticles(ctx *gin.Context) {
	search := search.ArticleElasticSearch{}
	searchQuery := ctx.Query("query")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	response, err := search.SearchArticles(searchQuery, pageNo, pageSize)
	if err != nil {
		buildArtSrchErrorResponse(err, ctx)
		return
	}
	buildArtSrchSuccessResponse(response, ctx)
}

func buildArtSrchErrorResponse(err error, ctx *gin.Context) {
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

func buildArtSrchSuccessResponse(articles []*domain.ArticleSearchType, ctx *gin.Context) {
	httpStatus := http.StatusOK
	res := response.ArticleSearchResponse{
		BaseResponse: response.BaseResponse{
			Code:    http.StatusText(httpStatus),
			Message: fmt.Sprintf(constants.SUCCESS_SEARCH_MESSAGE, len(articles))},
		Articles: articles,
	}
	ctx.JSON(httpStatus, res)
}
