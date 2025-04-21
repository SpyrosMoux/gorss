package models

import (
	"time"
)

type Article struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	ExternalId string    `json:"externalId" gorm:"not null;unique"`
	Title      string    `json:"title" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null"`
	Link       string    `json:"link" gorm:"not null;unique"`
	FeedID     string    `json:"FeedID" gorm:"index"`
	Feed       Feed      `gorm:"foreignKey:FeedID;references:Id"`
	ImageUrl   string    `json:"imageUrl"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
