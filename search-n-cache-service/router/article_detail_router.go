package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/service"
)

func SetupArticleDetailRouter(engine *gin.Engine) {

	engine.GET("/article/:id", func(ctx *gin.Context) {

		id := ctx.Params.ByName("id")
		article := service.GetArticleDetails(id);

		ctx.JSON(http.StatusOK, gin.H{
			"code":    "OK",
			"message": "Article Detail Response",
			"article": article,
		})

		/*ctx.JSON(200, gin.H{
			"code":    "OK",
			"message": "Article Detail Response",
		})*/
	})
}
