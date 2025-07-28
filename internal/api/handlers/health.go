package handlers

import (
	"net/http"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/pkg/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthHandler struct {
	logger *zap.Logger
}

func NewHealthHandler(logger *zap.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

func (h *HealthHandler) HealthCheck(ctx *gin.Context) {
	db := database.GetDB()
	var dbConnected string
	psqlDB, _ := db.DB()
	err := psqlDB.Ping()
	if err != nil {
		dbConnected = "disconnected"
	} else {
		dbConnected = "connected"
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "database": dbConnected, "timestamp": time.Now().UTC()})

}
