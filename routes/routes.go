package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/article"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/feed"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.Default())

	articleRepo := article.NewRepository(db.Conn)
	articleService := article.NewArticleService(articleRepo)
	articleHandler := article.NewArticleHandler(articleService)

	feedRepo := feed.NewFeedRepository(db.Conn)
	feedService := feed.NewFeedService(feedRepo, articleService)
	feedHandler := feed.NewFeedHandler(feedService)

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/healthz", HealthHandler)

		feedGroup := apiGroup.Group("/feeds")
		{
			feedGroup.POST("", feedHandler.HandleAddFeed)
			feedGroup.GET("", feedHandler.HandleGetAllFeeds)
		}

		articleGroup := apiGroup.Group("/articles")
		{
			articleGroup.GET("/latest", articleHandler.HandleGetLatestArticles)
			articleGroup.GET("/:feedId", articleHandler.HandleGetAllArticlesByFeedId)
		}
	}

	return router
}

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "I'm Alive!"})
}
