package repositories

import (
	"github.com/spyrosmoux/gorss/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateMany(articles *[]models.Article) error
}

type articleRepository struct {
	dbConn *gorm.DB
}

func NewArticleRepository(dbConn *gorm.DB) ArticleRepository {
	return &articleRepository{
		dbConn: dbConn,
	}
}

func (articleRepository articleRepository) CreateMany(articles *[]models.Article) error {
	result := articleRepository.dbConn.Create(&articles)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
