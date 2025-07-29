package handlers

import (
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArtistHandler struct {
	artistService *services.ArtistService
	logger        *zap.Logger
}

func NewArtistHandler(artistService *services.ArtistService, logger *zap.Logger) *ArtistHandler {
	return &ArtistHandler{
		artistService: artistService,
		logger:        logger,
	}
}

func (h *ArtistHandler) CreateArtist(ctx *gin.Context) {
	var req models.CreateArtistRequest
	id := ctx.Value("user_id").(int)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INVALID_JSON",
				Message: "Invalid JSON format",
				Details: err.Error(),
			},
		})
		return
	}
	response, err := h.artistService.CreateArtist(ctx.Request.Context(), &req, id)
	if err != nil {

		if appErr, ok := err.(*util.AppError); ok {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
				Status: false,
				Error:  appErr,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
				Details: err,
			},
		})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
