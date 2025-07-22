package handlers

import (
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "server running",
	})
}
