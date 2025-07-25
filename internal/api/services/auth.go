package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	customerror "github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
	logger      *slog.Logger
}

func NewAuthHandler(authService *services.AuthService, logger *slog.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invlaid request in registration request", slog.String("error", err.Error()))
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
	h.logger.Info("Registration attempt", slog.String("email", req.Username), slog.String("username", req.Email))

	response, err := h.authService.Register(ctx.Request.Context(), &req)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) handleError(ctx *gin.Context, err error) {
	if appError, ok := err.(*customerror.AppError); ok {
		ctx.JSON(appError.StatusCode, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    appError.Code,
				Message: appError.Message,
			},
		})
		return
	}
	h.logger.Error("Unhandled error in auth handler", slog.String("error", err.Error()))
	ctx.JSON(http.StatusInternalServerError, &util.ErrorData{
		Code:    "INTERNAL_ERROR",
		Message: "Internal server error",
	})
}
