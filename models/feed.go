package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feed struct {
	Id         string `json:"id" gorm:"primaryKey"`
	ExternalId string `json:"externalId" gorm:"unique"`
	Name       string `json:"name" gorm:"not null"`
	Link       string `json:"link" gorm:"not null"`
}

func (f *Feed) BeforeSave(tx *gorm.DB) error {
	if f.Id == "" {
		f.Id = uuid.NewString()
	}

	return nil
}
