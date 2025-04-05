package main

import (
  "log/slog"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/spyrosmoux/gorss/db"
  "github.com/spyrosmoux/gorss/models"
  "github.com/spyrosmoux/gorss/routes"
)

var router *gin.Engine

func init() {
  err := db.Connect()
  if err != nil {
    slog.Error(err.Error())
    os.Exit(1)
  }

  router = routes.SetupRouter()
}

type Feeder interface {
  Fetch(feed string) ([]byte, error)
  Parse([]byte) (models.Feed, []models.Article, error)
}

func main() {
  slog.Info("started server on", "port", "8080")
  err := router.Run(":8080")
  if err != nil {
    slog.Error("failed to start server", "port", "8080", "err", err)
    os.Exit(1)
  }
}
