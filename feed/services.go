package feed

import (
	"log/slog"

	"github.com/mmcdole/gofeed"
	"github.com/spyrosmoux/gorss/article"
	"github.com/spyrosmoux/gorss/models"
)

type FeedService interface {
	RegisterFeed(feedUrl string) error
	GetAllFeeds() ([]models.Feed, error)
}

type feedService struct {
	feedRepository FeedRepository
	feedParser     *gofeed.Parser
	articleService article.ArticleService
}

func NewFeedService(feedRepo FeedRepository, articleService article.ArticleService) FeedService {
	return &feedService{
		feedRepository: feedRepo,
		articleService: articleService,
		feedParser:     gofeed.NewParser(),
	}
}

func (feedService feedService) RegisterFeed(feedUrl string) error {
	goFeed, err := feedService.feedParser.ParseURL(feedUrl)
	if err != nil {
		slog.Error("failed to parse feed", "feedUrl", feedUrl, "err", err)
		return err
	}

	feed := FeedFromGoFeed(goFeed)

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
