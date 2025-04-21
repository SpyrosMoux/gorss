package article

import (
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/spyrosmoux/gorss/models"
)

func ArticleFromGoFeedItem(item *gofeed.Item) models.Article {
	var imageUrl string
	if item.Image != nil {
		imageUrl = item.Image.URL
	}
	return models.Article{
		Id:         uuid.NewString(),
		ExternalId: item.GUID,
		Title:      item.Title,
		Content:    item.Content,
		Link:       item.Link,
		ImageUrl:   imageUrl,
	}
}
