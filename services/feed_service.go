package services

import (
	"github.com/SpyrosMoux/gorss/models"
	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type FeedService interface {
	RegisterFeed(feedUrl string) error
	GetAllFeeds() ([]models.Feed, error)
}

type feedService struct {
	slogger        *zap.SugaredLogger
	feedRepository repositories.FeedRepository
	feedParser     *gofeed.Parser
	articleService ArticleService
}

func NewFeedService(slogger *zap.SugaredLogger, feedRepo repositories.FeedRepository, articleService ArticleService) FeedService {
	return &feedService{
		slogger:        slogger,
		feedRepository: feedRepo,
		articleService: articleService,
		feedParser:     gofeed.NewParser(),
	}
}

func (feedService feedService) RegisterFeed(feedUrl string) error {
	goFeed, err := feedService.feedParser.ParseURL(feedUrl)
	if err != nil {
		feedService.slogger.Errorw("failed to parse feed", "feedUrl", feedUrl, "err", err)
		return err
	}

	feed := feedFromGoFeed(goFeed)

	savedFeed, err := feedService.feedRepository.Create(feed)
	if err != nil {
		feedService.slogger.Errorw("failed to register", "feed", feed.Name, "err", err)
		return err
	}

	err = feedService.articleService.SyncArticlesByFeed(savedFeed)
	if err != nil {
		feedService.slogger.Errorw("failed to sync first-time articles", "feed", savedFeed.Name, "err", err)
		return err
	}

	feedService.slogger.Infow("registered a new feed", "feedUrl", savedFeed.Link)
	return nil
}

func (feedService feedService) GetAllFeeds() ([]models.Feed, error) {
	feeds, err := feedService.feedRepository.FindAll()
	if err != nil {
		feedService.slogger.Errorw("failed to get all feeds", "err", err)
		return nil, err
	}

	return feeds, nil
}

func feedFromGoFeed(goFeed *gofeed.Feed) models.Feed {
	return models.Feed{
		Id:       uuid.NewString(),
		Name:     goFeed.Title,
		Link:     goFeed.FeedLink,
		Articles: []models.Article{},
	}
}
