package models

import (
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"time"
)

type Article struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	ExternalId string    `json:"externalId" gorm:"not null"`
	Title      string    `json:"title" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null"`
	Link       string    `json:"link" gorm:"not null"`
	FeedID     string    `json:"FeedID" gorm:"index"`
	Feed       Feed      `gorm:"foreignKey:FeedID;references:Id"`
	ImageUrl   string    `json:"imageUrl"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func ArticleFromGoFeedItem(item *gofeed.Item) Article {
	var imageUrl string
	if item.Image != nil {
		imageUrl = item.Image.URL
	}
	return Article{
		Id:         uuid.NewString(),
		ExternalId: item.GUID,
		Title:      item.Title,
		Content:    item.Content,
		Link:       item.Link,
		ImageUrl:   imageUrl,
	}
}
