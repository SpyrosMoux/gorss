package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/services"
	"net/http"
)

type ArticleHandler interface {
	HandleGetLatestArticles(ctx *gin.Context)
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
