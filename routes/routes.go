package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/handlers"
	"github.com/spyrosmoux/gorss/repositories"
	"github.com/spyrosmoux/gorss/services"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.Default())

	articleRepo := repositories.NewArticleRepository(db.Conn)
	articleService := services.NewArticleService(articleRepo)
	articleHandler := handlers.NewArticleHandler(articleService)

	feedRepo := repositories.NewFeedRepository(db.Conn)
	feedService := services.NewFeedService(feedRepo, articleService)
	feedHandler := handlers.NewFeedHandler(feedService)

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/healthz", handlers.HealthHandler)

		feedGroup := apiGroup.Group("/feeds")
		{
			feedGroup.POST("", feedHandler.HandleAddFeed)
			feedGroup.GET("", feedHandler.HandleGetAllFeeds)
		}

		articleGroup := apiGroup.Group("/articles")
		{
			articleGroup.GET("/latest", articleHandler.HandleGetLatestArticles)
		}
	}

	return router
}
