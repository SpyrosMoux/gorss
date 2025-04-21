package feed

import (
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/spyrosmoux/gorss/models"
)

func FeedFromGoFeed(goFeed *gofeed.Feed) models.Feed {
	return models.Feed{
		Id:       uuid.NewString(),
		Name:     goFeed.Title,
		Link:     goFeed.FeedLink,
		Articles: []models.Article{},
	}
}
