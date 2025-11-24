package services

import (
	"fmt"
	"log/slog"

	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/mmcdole/gofeed"
)

type SchedulerService interface {
	SyncArticlesAllFeeds() error
}

type schedulerService struct {
	feedParser     *gofeed.Parser
	feedRepository repositories.FeedRepository
	articleService ArticleService
}

func NewSchedulerService(feedRepo repositories.FeedRepository, articleService ArticleService) SchedulerService {
	return &schedulerService{
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
			slog.Error("failed to sync feed", "feed", feed.Name, "error", err)
			syncErrors = append(syncErrors, err)
		}
	}

	if len(syncErrors) > 0 {
		return fmt.Errorf("failed to sync %d feeds", len(syncErrors))
	}

	return nil
}
