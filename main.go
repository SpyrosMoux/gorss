package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/spyrosmoux/gorss/atom"
	"github.com/spyrosmoux/gorss/internal/models"
)

type Feeder interface {
	Fetch(feed string) ([]byte, error)
	Parse([]byte) (models.Feed, []models.Article, error)
}

func main() {
	aggregator := atom.NewAtomAggregator()

	atomFeed, err := aggregator.Fetch("https://world.hey.com/dhh/feed.atom")
	if err != nil {
		os.Exit(1)
	}

	feed, articles, err := aggregator.Parse(atomFeed)
	if err != nil {
		slog.Error("failed to parse atom feed", "err", err)
		os.Exit(1)
	}

	slog.Info("fetched", "feed", feed, "articles", articles)

	slog.Info("started server on", "port", "8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("failed to start server", "port", "8080", "err", err)
		os.Exit(1)
	}
}
