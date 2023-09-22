package router

import (
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/service"
)

func SetupArticleDetailRouter(engine *gin.Engine) {
	// configure REST methods
	engine.POST("/article", service.SaveArticle)
	engine.GET("/article/:id", service.GetArticle)
	engine.PUT("/article", service.UpdateArticle)
	engine.DELETE("/article/:id", service.DeleteArticle)
}
