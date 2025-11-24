package repositories

import (
	"github.com/SpyrosMoux/gorss/models"
	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(feed models.Feed) (models.Feed, error)
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

func (feedRepo feedRepo) Create(feed models.Feed) (models.Feed, error) {
	result := feedRepo.dbConn.Create(&feed)
	if result.Error != nil {
		return models.Feed{}, result.Error
	}

	return feed, nil
}

func (feedRepo feedRepo) FindAll() ([]models.Feed, error) {
	var feeds []models.Feed
	result := feedRepo.dbConn.Find(&feeds)
	if result.Error != nil {
		return nil, result.Error
	}

	return feeds, nil
}
