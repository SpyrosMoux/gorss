package article

import (
	"fmt"

	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleRepository interface {
	CreateMany(articles []*models.Article) error
	FindAllByDate(orderDirection db.OrderDirection) ([]*models.Article, error)
	FindAllByFeedId(feedId string) ([]*models.Article, error)
}

type articleRepository struct {
	dbConn *gorm.DB
}

func NewRepository(dbConn *gorm.DB) ArticleRepository {
	return &articleRepository{
		dbConn: dbConn,
	}
}

// CreateMany inserts a batch of articles skipping duplicates
func (articleRepository articleRepository) CreateMany(articles []*models.Article) error {
	result := articleRepository.dbConn.Clauses(clause.OnConflict{DoNothing: true}).Create(&articles)
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

func (articleRepository articleRepository) FindAllByFeedId(feedId string) ([]*models.Article, error) {
	var articles []*models.Article
	result := articleRepository.dbConn.Find(&articles).Where("feedId = ?1", feedId)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}
