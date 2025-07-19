package handlers

import (
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "server running",
	})
}
