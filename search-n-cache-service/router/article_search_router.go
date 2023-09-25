package router

import (
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/service"
)

func SetupArticleListRouter(engine *gin.Engine) {
	// configure REST methods
	engine.GET("/article-search/list", service.GetArticleList)
}
