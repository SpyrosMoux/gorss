package article

import (
	"log/slog"

	"github.com/mmcdole/gofeed"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/models"
)

type ArticleService interface {
	SyncArticlesByFeed(feed models.Feed) error
	GetLatestArticles() ([]*ArticleDto, error)
	GetAllArticlesByFeedId(feedId string) ([]*ArticleDto, error)
}

type articleService struct {
	articleRepo ArticleRepository
	feedParser  *gofeed.Parser
}

func NewArticleService(articleRepo ArticleRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
		feedParser:  gofeed.NewParser(),
	}
}

func (articleService articleService) SyncArticlesByFeed(feed models.Feed) error {
	feedData, err := articleService.feedParser.ParseURL(feed.Link)
	if err != nil {
		slog.Error("failed to parse feed", "feed", feed.Id, "err", err)
		return err
	}

	var articles []*models.Article
	for _, item := range feedData.Items {
		article := ArticleFromGoFeedItem(item)
		article.FeedID = feed.Id
		article.Feed = feed
		articles = append(articles, &article)
	}

	err = articleService.articleRepo.CreateMany(articles)
	if err != nil {
		slog.Error("failed to sync articles", "feed", feed.Id, "err", err)
		return err
	}

	return nil
}

func (articleService articleService) GetLatestArticles() ([]*ArticleDto, error) {
	articles, err := articleService.articleRepo.FindAllByDate(db.ORDER_DESCENDING)
	if err != nil {
		slog.Error("failed to get latest articles", "err", err)
		return nil, err
	}

	var articleDtos []*ArticleDto
	for _, article := range articles {
		articleDto := articleToArticleDto(*article)
		articleDtos = append(articleDtos, &articleDto)
	}

	return articleDtos, nil
}

func (articleService articleService) GetAllArticlesByFeedId(feedId string) ([]*ArticleDto, error) {
	articles, err := articleService.articleRepo.FindAllByFeedId(feedId)
	if err != nil {
		slog.Error("failed to get articles by feed id", "feedId", feedId, "err", err)
		return nil, err
	}

	var articleDtos []*ArticleDto
	for _, article := range articles {
		articleDto := articleToArticleDto(*article)
		articleDtos = append(articleDtos, &articleDto)
	}

	return articleDtos, nil
}

func articleToArticleDto(article models.Article) ArticleDto {
	return ArticleDto{
		Id:       article.Id,
		Title:    article.Title,
		Content:  article.Content,
		Link:     article.Link,
		FeedId:   article.FeedID,
		ImageUrl: article.ImageUrl,
		Date:     article.UpdatedAt,
	}
}
