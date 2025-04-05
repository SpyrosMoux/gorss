package atom

import (
	"encoding/xml"
	"io"
	"log/slog"
	"net/http"

	"github.com/spyrosmoux/gorss/models"
)

type AtomAggregator struct{}

func NewAtomAggregator() AtomAggregator {
	return AtomAggregator{}
}

func (aggr AtomAggregator) Fetch(feed string) ([]byte, error) {
	resp, err := http.Get(feed)
	if err != nil {
		slog.Error("failed to fetch feed", "feed", feed, "err", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body", "feed", feed, "err", err)
		return []byte{}, err
	}

	return body, nil
}

func (aggr AtomAggregator) Parse(body []byte) (models.Feed, []models.Article, error) {
	var atomFeed AtomFeed
	err := xml.Unmarshal(body, &atomFeed)
	if err != nil {
		slog.Error("failed to unmarshal feed", "err", err)
		return models.Feed{}, []models.Article{}, nil
	}

	feed := models.Feed{
		ExternalId: atomFeed.ID,
		Name:       atomFeed.Title,
		Link:       atomFeed.Link[0].Href,
	}

	var articles []models.Article
	for _, entry := range atomFeed.Entries {
		article := models.Article{
			ExternalId: entry.ID,
			Title:      entry.Title,
			Content:    entry.Content,
			Link:       entry.Link.Href,
		}
		articles = append(articles, article)
	}

	return feed, articles, nil
}
