package scheduler

import (
	"github.com/mmcdole/gofeed"
	"github.com/spyrosmoux/gorss/article"
	"github.com/spyrosmoux/gorss/feed"
)

type SchedulerService interface {
	SyncArticlesAllFeeds() error
}

type schedulerService struct {
	feedParser     *gofeed.Parser
	feedRepository feed.FeedRepository
	articleService article.ArticleService
}

func NewSchedulerService(feedRepo feed.FeedRepository, articleService article.ArticleService) SchedulerService {
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

	for _, feed := range feeds {
		err := schedulerService.articleService.SyncArticlesByFeed(feed)
		if err != nil {
			return err
		}
	}

	return nil
}
