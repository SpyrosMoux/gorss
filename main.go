package main

import (
	"log/slog"
	"os"
	"time"

	v1 "github.com/SpyrosMoux/gorss/api/v1"
	"github.com/SpyrosMoux/gorss/db"
	"github.com/SpyrosMoux/gorss/env"
	"github.com/SpyrosMoux/gorss/models"
	"github.com/SpyrosMoux/gorss/repositories"
	"github.com/SpyrosMoux/gorss/services"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	apiPort     string
	dbHost      string
	dbPort      string
	dbUser      string
	dbPass      string
	dbName      string
	environment string
	router      *gin.Engine
	scheduler   *models.Scheduler
	slogger     *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger = logger.Sugar()

	apiPort = env.LoadEnvVariable("API_PORT")
	dbHost = env.LoadEnvVariable("DB_HOST")
	dbPort = env.LoadEnvVariable("DB_PORT")
	dbUser = env.LoadEnvVariable("DB_USER")
	dbPass = env.LoadEnvVariable("DB_PASS")
	dbName = env.LoadEnvVariable("DB_NAME")
	environment = env.LoadEnvVariable("ENVIRONMENT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	err := db.Init(dsn, "gorss", &models.Feed{}, &models.Article{})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	articleRepository := repositories.NewRepository(db.Conn)
	articleService := services.NewArticleService(slogger, articleRepository)
	articleHandlerV1 := v1.NewArticleHandler(articleService)

	feedRepository := repositories.NewFeedRepository(db.Conn)
	feedService := services.NewFeedService(slogger, feedRepository, articleService)
	feedHandlerV1 := v1.NewFeedHandler(feedService)

	router = gin.New()
	router.Use(ginzap.Ginzap(slogger.Desugar(), time.RFC3339, true))

	if environment == "UAT" || environment == "DEV" {
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000", "http://localhost:5173", "https://uat-gorss.spyrosmoux.com"},
		}))
	}

	apiV1 := router.Group("/api/v1")
	v1.RegisterV1Routes(apiV1, articleHandlerV1, feedHandlerV1)

	setupScheduler()
}

func main() {
	scheduler.Start()
	defer scheduler.Stop()

	slogger.Infow("started server", "port", apiPort)
	err := router.Run(":" + apiPort)
	if err != nil {
		slogger.Fatalw("failed to start server", "port", apiPort, "err", err)
	}
}

func setupScheduler() {
	articleRepo := repositories.NewRepository(db.Conn)
	articleService := services.NewArticleService(slogger, articleRepo)
	feedRepo := repositories.NewFeedRepository(db.Conn)
	schedulerService := services.NewSchedulerService(slogger, feedRepo, articleService)
	scheduler = models.NewScheduler(slogger)
	scheduler.AddTask("SyncAllFeeds", time.Hour, schedulerService.SyncArticlesAllFeeds)
}
