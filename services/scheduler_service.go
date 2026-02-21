package services

import (
	"fmt"

	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type SchedulerService interface {
	SyncArticlesAllFeeds() error
}

type schedulerService struct {
	slogger        *zap.SugaredLogger
	feedParser     *gofeed.Parser
	feedRepository repositories.FeedRepository
	articleService ArticleService
}

func NewSchedulerService(slogger *zap.SugaredLogger, feedRepo repositories.FeedRepository, articleService ArticleService) SchedulerService {
	return &schedulerService{
		slogger:        slogger,
		feedParser:     gofeed.NewParser(),
		feedRepository: feedRepo,
		articleService: articleService,
	}
}

func (schedulerService schedulerService) SyncArticlesAllFeeds() error {
	feeds, err := schedulerService.feedRepository.FindAll()
	if err != nil {
		return err
	}

	var syncErrors []error
	for _, feed := range feeds {
		err := schedulerService.articleService.SyncArticlesByFeed(feed)
		if err != nil {
			schedulerService.slogger.Errorw("failed to sync feed", "feed", feed.Name, "err", err)
			syncErrors = append(syncErrors, err)
		}
	}

	if len(syncErrors) > 0 {
		return fmt.Errorf("failed to sync %d feeds", len(syncErrors))
	}

	return nil
}
