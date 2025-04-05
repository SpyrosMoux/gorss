package services

import (
	"github.com/spyrosmoux/gorss/atom"
	"github.com/spyrosmoux/gorss/models"
	"github.com/spyrosmoux/gorss/repositories"
	"log/slog"
)

type ArticleService interface {
	SyncArticlesByFeed(feed models.Feed) error
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

	err = articleService.articleRepo.CreateMany(&articles)
	if err != nil {
		slog.Error("failed to sync articles", "feed", feed.Id, "err", err)
		return err
	}

	return nil
}
