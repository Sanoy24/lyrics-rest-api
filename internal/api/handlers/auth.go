package handlers

import (
	"errors"
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	customerror "github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService *services.AuthService
	logger      *zap.Logger
}

func NewAuthHandler(authService *services.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// h.logger.Warn("Invlaid request in registration request", zap.String("error", err.Error()))
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
	// h.logger.Info("Registration attempt", zap.String("email", req.Username), zap.String("username", req.Email))

	response, err := h.authService.Register(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server eerror",
				Details: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req models.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Info("Invalid request in login request", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, customerror.ErrInvalidCredentials)
		return
	}
	response, err := h.authService.Login(ctx.Request.Context(), &req)
	if errors.Is(err, customerror.ErrInvalidCredentials) {
		ctx.JSON(http.StatusUnauthorized, customerror.ErrInvalidCredentials)
		return
	}
	if errors.Is(err, customerror.ErrInternalServer) {
		ctx.JSON(http.StatusInternalServerError, customerror.ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, response)

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
	h.logger.Error("Unhandled error in auth handler", zap.String("error", err.Error()))
	ctx.JSON(http.StatusInternalServerError, &util.ErrorData{
		Code:    "INTERNAL_ERROR",
		Message: "Internal server error",
	})
}
