package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/constants"
	"github.com/search-n-cache/search-n-cache-service/domain"
	"github.com/search-n-cache/search-n-cache-service/error"
	"github.com/search-n-cache/search-n-cache-service/response"
	"github.com/search-n-cache/search-n-cache-service/search"
)

// Get Article List
//
//	@Summary      Get article List
//	@Description  Get Article List Information
//	@Tags         Article Search
//	@Accept       json
//	@Produce      json
//	@Param 		  pageNo     query      int  false  "Page No"
//	@Param 		  pageSize   query      int  false  "Page Size"
//	@Success      200  {object}  response.ArticleDetailResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article-search/list [get]
func GetArticleList(ctx *gin.Context) {
	search := search.ArticleElasticSearch{}
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	response, err := search.GetArticleList(pageNo, pageSize)
	if err != nil {
		error.HandleError(err, ctx)
		return
	}
	buildArtSrchSuccessResponse(response, ctx)
}

// Search Articles 
//
//	@Summary      Search articles
//	@Description  Search Article Information
//	@Tags         Article Search
//	@Accept       json
//	@Produce      json
//	@Param 		  query      query      string  true "Search Query"
//	@Param 		  pageNo     query      int     false  "Page No"
//	@Param 		  pageSize   query      int     false  "Page Size"
//	@Success      200  {object}  response.ArticleDetailResponse
//	@Failure      500  {object}  response.ErrorResponse
//	@Router       /article-search/search [get]
func SearchArticles(ctx *gin.Context) {
	search := search.ArticleElasticSearch{}
	searchQuery := ctx.Query("query")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	response, err := search.SearchArticles(searchQuery, pageNo, pageSize)
	if err != nil {
		error.HandleError(err, ctx)
		return
	}
	buildArtSrchSuccessResponse(response, ctx)
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
