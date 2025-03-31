package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feed struct {
	Id   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Article struct {
	Id      string `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Link    string `json:"link"`
}

func (f *Feed) BeforeSave(tx *gorm.DB) error {
	if f.Id == "" {
		f.Id = uuid.NewString()
	}

	return nil
}

func (a *Article) BeforeSave(tx *gorm.DB) error {
	if a.Id == "" {
		a.Id = uuid.NewString()
	}

	return nil
}
