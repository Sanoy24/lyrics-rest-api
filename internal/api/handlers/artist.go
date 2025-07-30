package handlers

import (
	"net/http"
	"strconv"

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

func (h *ArtistHandler) GetAllArtists(ctx *gin.Context) {
	response, err := h.artistService.GetAllArtists(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ArtistHandler) GetArtistByID(ctx *gin.Context) {
	id := ctx.Param("id")
	artistID, _ := strconv.Atoi(id)
	response, err := h.artistService.GetArtistByID(ctx.Request.Context(), artistID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorResponse{
				Status: false,
				Error:  "Internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, response)

}

func (h *ArtistHandler) UpdateArtist(ctx *gin.Context) {
	var req models.UpdateArtistRequest
	id := ctx.Param("id")
	artistID, _ := strconv.Atoi(id)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_SERVER",
				Message: "Internal server error",
				Details: err.Error(),
			},
		})
		return
	}
	response, err := h.artistService.UpdateArtist(ctx.Request.Context(), artistID, &req)
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
				Code:    "INTERNAL_SERVER",
				Message: "Internal server error",
				Details: err,
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
func (h *ArtistHandler) DeleteArtist(ctx *gin.Context) {
	id := ctx.Param("id")
	artistID, _ := strconv.Atoi(id)
	response, err := h.artistService.DeleteArtist(ctx.Request.Context(), artistID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_SERVER",
				Message: "Internal server error",
				Details: err.Error(),
			},
		})
		return

	}
	ctx.JSON(http.StatusOK, response)
}
