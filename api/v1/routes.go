package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(v1Group *gin.RouterGroup, articleHandler ArticleHandler, feedHandler FeedHandler) {
	v1Group.GET("/healthz", healthHandler)

	// register article handlers
	articleGroup := v1Group.Group("/articles")
	{
		articleGroup.GET("/latest", articleHandler.HandleGetLatestArticles)
		articleGroup.GET("/:feedId", articleHandler.HandleGetAllArticlesByFeedId)
	}

	feedGroup := v1Group.Group("/feeds")
	{
		feedGroup.POST("", feedHandler.HandleAddFeed)
		feedGroup.GET("", feedHandler.HandleGetAllFeeds)
	}
}

func healthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "I'm Alive!"})
}
