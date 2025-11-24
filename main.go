package main

import (
	"log/slog"
	"os"
	"time"

	v1 "github.com/SpyrosMoux/gorss/api/v1"
	"github.com/SpyrosMoux/gorss/db"
	"github.com/SpyrosMoux/gorss/models"
	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/SpyrosMoux/gorss/services"
	"github.com/SpyrosMoux/helpers/env"
	"github.com/gin-gonic/gin"
)

var (
	apiPort   string
	dbHost    string
	dbPort    string
	dbUser    string
	dbPass    string
	dbName    string
	router    *gin.Engine
	scheduler *models.Scheduler
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

	articleRepository := repositories.NewRepository(db.Conn)
	articleService := services.NewArticleService(articleRepository)
	articleHandlerV1 := v1.NewArticleHandler(articleService)

	feedRepository := repositories.NewFeedRepository(db.Conn)
	feedService := services.NewFeedService(feedRepository, articleService)
	feedHandlerV1 := v1.NewFeedHandler(feedService)

	router = gin.Default()
	apiV1 := router.Group("/api/v1")
	v1.RegisterV1Routes(apiV1, articleHandlerV1, feedHandlerV1)

	setupScheduler()
}

func main() {
	scheduler.Start()
	defer scheduler.Stop()

	slog.Info("started server on", "port", apiPort)
	err := router.Run(":" + apiPort)
	if err != nil {
		slog.Error("failed to start server", "port", apiPort, "err", err)
		os.Exit(1)
	}
}

func setupScheduler() {
	articleRepo := repositories.NewRepository(db.Conn)
	articleService := services.NewArticleService(articleRepo)
	feedRepo := repositories.NewFeedRepository(db.Conn)
	schedulerService := services.NewSchedulerService(feedRepo, articleService)
	scheduler = models.NewScheduler()
	scheduler.AddTask("SyncAllFeeds", 10*time.Second, schedulerService.SyncArticlesAllFeeds)
}
