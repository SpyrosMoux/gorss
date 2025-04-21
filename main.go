package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/article"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/feed"
	"github.com/spyrosmoux/gorss/models"
	"github.com/spyrosmoux/gorss/routes"
	"github.com/spyrosmoux/gorss/scheduler"
)

var router *gin.Engine
var sch *scheduler.Scheduler

func init() {
	err := db.Connect(&models.Feed{}, &models.Article{})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	router = routes.SetupRouter()
	setupScheduler()
}

func main() {
	sch.Start()
	defer sch.Stop()

	slog.Info("started server on", "port", "8080")
	err := router.Run(":8080")
	if err != nil {
		slog.Error("failed to start server", "port", "8080", "err", err)
		os.Exit(1)
	}
}

func setupScheduler() {
	articleRepo := article.NewRepository(db.Conn)
	articleService := article.NewArticleService(articleRepo)
	feedRepo := feed.NewFeedRepository(db.Conn)
	schedulerService := scheduler.NewSchedulerService(feedRepo, articleService)
	sch = scheduler.NewScheduler()
	sch.AddTask("SyncAllFeeds", time.Hour, schedulerService.SyncArticlesAllFeeds)
}
