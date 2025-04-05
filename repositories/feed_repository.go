package repositories

import (
	"github.com/spyrosmoux/gorss/models"
	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(feed models.Feed) error
	FindAll() ([]models.Feed, error)
}

type feedRepo struct {
	dbConn *gorm.DB
}

func NewFeedRepository(dbConn *gorm.DB) FeedRepository {
	return &feedRepo{
		dbConn: dbConn,
	}
}

func (feedRepo feedRepo) Create(feed models.Feed) error {
	result := feedRepo.dbConn.Create(&feed)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (feedRepo feedRepo) FindAll() ([]models.Feed, error) {
	var feeds []models.Feed
	result := feedRepo.dbConn.Find(&feeds)
	if result.Error != nil {
		return nil, result.Error
	}

	return feeds, nil
}
