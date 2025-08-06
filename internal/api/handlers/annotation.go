package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	customerror "github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
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

func (a *AnnotationHandler) UpdateAnnotation(ctx *gin.Context) {
	var req models.UpdateAnnotationRequest
	songID, _ := strconv.Atoi(ctx.Param("song_id"))
	annotationID, _ := strconv.Atoi(ctx.Param("annotation_id"))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		a.logger.Error("failed to bind json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
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
		return
	}
	if err := a.annotationService.UpdateAnnotation(ctx.Request.Context(), annotationID, songID, &req); err != nil {
		if errors.Is(err, customerror.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, &util.ErrorResponse{
				Status: false,
				Error: &util.ErrorData{
					Code:    "NOT_FOUND",
					Message: "Annotation Not Found",
					Details: err.Error(),
				},
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
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
		Message: "annotation updated successfully",
	})
}

func (a *AnnotationHandler) DeleteAnnotation(ctx *gin.Context) {
	annotationID, _ := strconv.Atoi(ctx.Param("id"))
	if err := a.annotationService.DeleteAnnotation(ctx.Request.Context(), annotationID); err != nil {
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
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
		Message: "annotation deleted successfully",
	})
}
