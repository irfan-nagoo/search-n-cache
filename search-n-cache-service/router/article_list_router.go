package router

import (
	"github.com/gin-gonic/gin"
)

func SetupArticleListRouter(engine *gin.Engine) {

	engine.GET("/article-list/list", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    "OK",
			"message": "Article List Response",
		})
	})
}
