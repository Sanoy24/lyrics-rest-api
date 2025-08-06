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

type AnnotationHandler struct {
	annotationService *services.AnnotationService
	logger            *zap.Logger
}

func NewAnnotationHandler(annotationService *services.AnnotationService, logger *zap.Logger) *AnnotationHandler {
	return &AnnotationHandler{
		annotationService: annotationService,
		logger:            logger,
	}
}

func (a *AnnotationHandler) CreateAnnotation(ctx *gin.Context) {
	var req models.CreateAnnotationRequest

	songId, _ := strconv.Atoi(ctx.Param("song_id"))
	userId := ctx.Value("user_id").(int)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		a.logger.Error("failed to bind json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INVALID_JSON",
				Message: "Invalid JSON format",
				Details: err.Error(),
			},
		},
		)
		return

	}

	if appErr := util.ValidateStruct(&req); appErr != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error:  appErr,
		})
		return
	}

	annotationData, err := a.annotationService.CreateAnnotation(ctx.Request.Context(), songId, userId, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
				Details: err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusCreated, &util.SuccessResponse{
		Status:  true,
		Message: "annotation created successfully",
		Data:    annotationData,
	})

}

func (a *AnnotationHandler) GetAnnotationsBySongID(ctx *gin.Context) {
	songID, _ := strconv.Atoi(ctx.Param("song_id"))
	annotations, err := a.annotationService.GetAnnotationsBySongID(ctx.Request.Context(), songID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
				Details: err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &util.SuccessResponse{
		Status:  true,
		Message: "Annotations fetched successfully",
		Data:    annotations,
	})

}
