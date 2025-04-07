package repositories

import (
	"fmt"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateMany(articles []*models.Article) error
	FindAllByDate(orderDirection db.OrderDirection) ([]*models.Article, error)
}

type articleRepository struct {
	dbConn *gorm.DB
}

func NewArticleRepository(dbConn *gorm.DB) ArticleRepository {
	return &articleRepository{
		dbConn: dbConn,
	}
}

func (articleRepository articleRepository) CreateMany(articles []*models.Article) error {
	result := articleRepository.dbConn.Create(&articles)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleRepository articleRepository) FindAllByDate(orderDirection db.OrderDirection) ([]*models.Article, error) {
	order := fmt.Sprintf("updated_at %s", orderDirection.String())

	var articles []*models.Article
	result := articleRepository.dbConn.Order(order).Limit(5).Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}
