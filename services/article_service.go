package services

import (
	"log/slog"

	"github.com/SpyrosMoux/gorss/db"
	"github.com/SpyrosMoux/gorss/models"
	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type ArticleService interface {
	SyncArticlesByFeed(feed models.Feed) error
	GetLatestArticles() ([]*models.ArticleDto, error)
	GetAllArticlesByFeedId(feedId string) ([]*models.ArticleDto, error)
}

type articleService struct {
	articleRepo repositories.ArticleRepository
	feedParser  *gofeed.Parser
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
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
		article := articleFromGoFeedItem(item)
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

func (articleService articleService) GetLatestArticles() ([]*models.ArticleDto, error) {
	articles, err := articleService.articleRepo.FindAllByDate(db.ORDER_DESCENDING)
	if err != nil {
		slog.Error("failed to get latest articles", "err", err)
		return nil, err
	}

	var articleDtos []*models.ArticleDto
	for _, article := range articles {
		articleDto := articleToArticleDto(*article)
		articleDtos = append(articleDtos, &articleDto)
	}

	return articleDtos, nil
}

func (articleService articleService) GetAllArticlesByFeedId(feedId string) ([]*models.ArticleDto, error) {
	articles, err := articleService.articleRepo.FindAllByFeedId(feedId)
	if err != nil {
		slog.Error("failed to get articles by feed id", "feedId", feedId, "err", err)
		return nil, err
	}

	var articleDtos []*models.ArticleDto
	for _, article := range articles {
		articleDto := articleToArticleDto(*article)
		articleDtos = append(articleDtos, &articleDto)
	}

	return articleDtos, nil
}

func articleToArticleDto(article models.Article) models.ArticleDto {
	return models.ArticleDto{
		Id:       article.Id,
		Title:    article.Title,
		Content:  article.Content,
		Link:     article.Link,
		FeedId:   article.FeedID,
		ImageUrl: article.ImageUrl,
		Date:     article.UpdatedAt,
	}
}

func articleFromGoFeedItem(item *gofeed.Item) models.Article {
	var imageUrl string
	if item.Image != nil {
		imageUrl = item.Image.URL
	}
	return models.Article{
		Id:         uuid.NewString(),
		ExternalId: item.GUID,
		Title:      item.Title,
		Content:    item.Content,
		Link:       item.Link,
		ImageUrl:   imageUrl,
	}
}
