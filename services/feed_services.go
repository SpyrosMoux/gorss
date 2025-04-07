package services

import (
	"log/slog"

	"github.com/spyrosmoux/gorss/atom"
	"github.com/spyrosmoux/gorss/models"
	"github.com/spyrosmoux/gorss/repositories"
)

type FeedService interface {
	RegisterFeed(feedUrl, feedType string) error
	GetAllFeeds() ([]models.Feed, error)
}

type feedService struct {
	feedRepository repositories.FeedRepository
	articleService ArticleService
}

func NewFeedService(feedRepo repositories.FeedRepository, articleService ArticleService) FeedService {
	return &feedService{
		feedRepository: feedRepo,
		articleService: articleService,
	}
}

func (feedService feedService) RegisterFeed(feedUrl, feedType string) error {
	feed, err := registerAtomFeed(feedUrl)
	if err != nil {
		return err
	}

	savedFeed, err := feedService.feedRepository.Create(feed)
	if err != nil {
		slog.Error("failed to register", "feed", feed.Name, "err", err)
		return err
	}

	err = feedService.articleService.SyncArticlesByFeed(savedFeed)
	if err != nil {
		slog.Error("failed to sync first-time articles", "feed", savedFeed.Name, "err", err)
		return err
	}

	slog.Info("registered a new feed", "feedUrl", savedFeed.Link)
	return nil
}

func (feedService feedService) GetAllFeeds() ([]models.Feed, error) {
	feeds, err := feedService.feedRepository.FindAll()
	if err != nil {
		slog.Error("failed to get all feeds", "err", err)
		return nil, err
	}

	return feeds, nil
}

func registerAtomFeed(feedUrl string) (models.Feed, error) {
	atomAggr := atom.NewAtomAggregator()
	atomFeed, err := atomAggr.Fetch(feedUrl)
	if err != nil {
		return models.Feed{}, err
	}

	feed, _, err := atomAggr.Parse(atomFeed)
	if err != nil {
		return models.Feed{}, err
	}

	feed.Link = feedUrl
	return feed, nil
}
