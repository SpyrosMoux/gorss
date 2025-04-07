package services

import (
	"github.com/spyrosmoux/gorss/atom"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/dto"
	"github.com/spyrosmoux/gorss/models"
	"github.com/spyrosmoux/gorss/repositories"
	"log/slog"
)

type ArticleService interface {
	SyncArticlesByFeed(feed models.Feed) error
	GetLatestArticles() ([]*dto.ArticleDto, error)
}

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

func (articleService articleService) SyncArticlesByFeed(feed models.Feed) error {
	atomAggr := atom.NewAtomAggregator()
	feedBytes, err := atomAggr.Fetch(feed.Link)
	if err != nil {
		slog.Error("failed to fetch feed", "feed", feed.Id, "err", err)
		return err
	}

	_, articles, err := atomAggr.Parse(feedBytes)
	if err != nil {
		slog.Error("failed to parse articles", "feed", feed.Id, "err", err)
		return err
	}

	for _, article := range articles {
		article.FeedID = feed.Id
		article.Feed = feed
	}

	err = articleService.articleRepo.CreateMany(articles)
	if err != nil {
		slog.Error("failed to sync articles", "feed", feed.Id, "err", err)
		return err
	}

	return nil
}

func (articleService articleService) GetLatestArticles() ([]*dto.ArticleDto, error) {
	articles, err := articleService.articleRepo.FindAllByDate(db.ORDER_DESCENDING)
	if err != nil {
		slog.Error("failed to get latest articles", "err", err)
		return nil, err
	}

	var articleDtos []*dto.ArticleDto
	for _, article := range articles {
		articleDto := articleToArticleDto(*article)
		articleDtos = append(articleDtos, &articleDto)
	}

	return articleDtos, nil
}

func articleToArticleDto(article models.Article) dto.ArticleDto {
	return dto.ArticleDto{
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
		Link:    article.Link,
		FeedId:  article.FeedID,
		Date:    article.UpdatedAt,
	}
}
