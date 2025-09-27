package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spyrosmoux/gorss/article"
	"github.com/spyrosmoux/gorss/db"
	"github.com/spyrosmoux/gorss/env"
	"github.com/spyrosmoux/gorss/feed"
	"github.com/spyrosmoux/gorss/models"
	"github.com/spyrosmoux/gorss/routes"
	"github.com/spyrosmoux/gorss/scheduler"
)

var (
	apiPort string
	dbHost  string
	dbPort  string
	dbUser  string
	dbPass  string
	dbName  string
	router  *gin.Engine
	sch     *scheduler.Scheduler
)

func init() {
	apiPort = env.LoadEnvVariable("API_PORT")
	dbHost = env.LoadEnvVariable("DB_HOST")
	dbPort = env.LoadEnvVariable("DB_PORT")
	dbUser = env.LoadEnvVariable("DB_USER")
	dbPass = env.LoadEnvVariable("DB_PASS")
	dbName = env.LoadEnvVariable("DB_NAME")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	err := db.Init(dsn, "gorss", &models.Feed{}, &models.Article{})
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

	slog.Info("started server on", "port", apiPort)
	err := router.Run(":" + apiPort)
	if err != nil {
		slog.Error("failed to start server", "port", apiPort, "err", err)
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
