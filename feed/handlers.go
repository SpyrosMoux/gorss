package feed

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedHandler interface {
	HandleAddFeed(ctx *gin.Context)
	HandleGetAllFeeds(ctx *gin.Context)
}

type feedHandler struct {
	feedService FeedService
}

func NewFeedHandler(feedService FeedService) FeedHandler {
	return &feedHandler{
		feedService: feedService,
	}
}

func (feedHandler feedHandler) HandleAddFeed(ctx *gin.Context) {
	var addFeedDto AddFeedDto
	err := ctx.ShouldBindJSON(&addFeedDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = feedHandler.feedService.RegisterFeed(addFeedDto.FeedUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (feedHandler feedHandler) HandleGetAllFeeds(ctx *gin.Context) {
	feeds, err := feedHandler.feedService.GetAllFeeds()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"feeds": feeds})
}
