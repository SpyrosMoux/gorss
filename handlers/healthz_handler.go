package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "I'm Alive!"})
}
