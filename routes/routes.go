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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:63342"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowCredentials: true,
	}))

	articleRepo := repositories.NewArticleRepository(db.Conn)
	articleService := services.NewArticleService(articleRepo)

	feedRepo := repositories.NewFeedRepository(db.Conn)
	feedService := services.NewFeedService(feedRepo, articleService)
	feedHandler := handlers.NewFeedHandler(feedService)

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/healthz", handlers.HealthHandler)

		feedGroup := apiGroup.Group("/feeds")
		{
			feedGroup.POST("", feedHandler.HandleAddFeed)
		}
	}

	return router
}
