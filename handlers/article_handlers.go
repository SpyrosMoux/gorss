package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/services"
	"net/http"
)

type ArticleHandler interface {
	HandleGetLatestArticles(ctx *gin.Context)
	HandleGetAllArticlesByFeedId(ctx *gin.Context)
}

type articleHandler struct {
	articleService services.ArticleService
}

func NewArticleHandler(articleService services.ArticleService) ArticleHandler {
	return &articleHandler{articleService: articleService}
}

func (a articleHandler) HandleGetLatestArticles(ctx *gin.Context) {
	articleDtos, err := a.articleService.GetLatestArticles()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"articles": articleDtos})
}

func (a articleHandler) HandleGetAllArticlesByFeedId(ctx *gin.Context) {
	feedId := ctx.Param("feedId")
	if feedId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "feedId cannot be empty"})
		return
	}

	articleDtos, err := a.articleService.GetAllArticlesByFeedId(feedId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"articles": articleDtos})
}
