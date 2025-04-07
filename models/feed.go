package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	Id         string `json:"id" gorm:"primaryKey"`
	ExternalId string `json:"externalId" gorm:"unique"`
	Name       string `json:"name" gorm:"not null"`
	Link       string `json:"link" gorm:"not null"`
	Articles   []Article
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (f *Feed) BeforeSave(tx *gorm.DB) error {
	if f.Id == "" {
		f.Id = uuid.NewString()
	}

	return nil
}
