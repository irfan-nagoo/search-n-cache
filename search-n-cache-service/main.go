package main

import (
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/config"
	"github.com/search-n-cache/search-n-cache-service/router"
	"github.com/search-n-cache/search-n-cache-service/search"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/search-n-cache/search-n-cache-service/docs"
)

func main() {
	// Setup env configuration
	config.SetEnvironmentConfg()

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	// Setup Article Router
	router.SetupArticleListRouter(engine)
	router.SetupArticleDetailRouter(engine)

	// Initialize DB connection
	component.DBConnection = component.InitializeDBConnection()

	// Initialize Search client (Default: ElasticSearch)
	component.ElasticTypedClient = component.InitializeElasticSearchClient()
	(&search.ArticleElasticSearch{}).CreateIndex()

	// Initialize Cache client (Default: Redis)
	component.RedisClient = component.InitializeRedisCacheClient()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engine.Run()
}
