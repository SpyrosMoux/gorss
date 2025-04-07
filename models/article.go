package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (a *Article) BeforeSave(tx *gorm.DB) error {
	if a.Id == "" {
		a.Id = uuid.NewString()
	}

	return nil
}
