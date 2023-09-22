package main

import (
	"github.com/gin-gonic/gin"
	"github.com/search-n-cache/search-n-cache-service/component"
	"github.com/search-n-cache/search-n-cache-service/config"
	"github.com/search-n-cache/search-n-cache-service/router"
)

func main() {
	// Setup env configuration
	config.SetEnvironmentConfg()

	engine := gin.Default()
	// Setup Article Router
	router.SetupArticleListRouter(engine)
	router.SetupArticleDetailRouter(engine)

	// Initialize DB connection
	component.DBConnection = component.InitializeDBConnection()

	engine.Run()
}
