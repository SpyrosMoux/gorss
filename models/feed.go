package models

import "time"

type Feed struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null;unique"`
	Link      string `json:"link" gorm:"not null;unique"`
	Articles  []Article
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
