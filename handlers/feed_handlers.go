package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/dto"
	"github.com/spyrosmoux/gorss/services"
)

type FeedHandler interface {
	HandleAddFeed(ctx *gin.Context)
}

type feedHandler struct {
	feedService services.FeedService
}

func NewFeedHandler(feedService services.FeedService) FeedHandler {
	return &feedHandler{
		feedService: feedService,
	}
}

func (feedHandler feedHandler) HandleAddFeed(ctx *gin.Context) {
	var addFeedDto dto.AddFeedDto
	err := ctx.ShouldBindJSON(&addFeedDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = feedHandler.feedService.RegisterFeed(addFeedDto.FeedUrl, addFeedDto.FeedType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
}
