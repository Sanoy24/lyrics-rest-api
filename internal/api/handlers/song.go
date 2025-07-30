package handlers

import (
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SongHandler struct {
	songService *services.SongService
	logger      *zap.Logger
}

func NewSongHandler(songService *services.SongService, logger *zap.Logger) *SongHandler {
	return &SongHandler{
		songService: songService,
		logger:      logger,
	}
}

func (h *SongHandler) CreateSong(ctx *gin.Context) {
	var req models.CreateSongRequest
	id := ctx.Value("user_id").(int)
	req.ContributorID = uint(id)
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
	if appErr := util.ValidateStruct(&req); appErr != nil {

		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error:  appErr,
		})
	}
	response, err := h.songService.CreateSong(ctx.Request.Context(), &req)
	if err != nil {

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
	ctx.JSON(http.StatusCreated, util.SuccessResponse{
		Status: true,
		Data:   response,
	})
}
