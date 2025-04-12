package models

import (
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"time"
)

type Feed struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Link      string `json:"link" gorm:"not null"`
	Articles  []Article
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FeedFromGoFeed(goFeed *gofeed.Feed) Feed {
	return Feed{
		Id:       uuid.NewString(),
		Name:     goFeed.Title,
		Link:     goFeed.FeedLink,
		Articles: []Article{},
	}
}
