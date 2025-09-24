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

type UserHandler struct {
	userService *services.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService *services.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}

}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Retrieve the details of the currently authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} util.SuccessResponse "User retrieved successfully"
// @Failure 500 {object} util.ErrorResponse "Internal server error"
// @Router /users/me [get]
func (u *UserHandler) GetCurrentUser(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		})
	}
	user, err := u.userService.GetCurrentUser(ctx.Request.Context(), userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		},
		)
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "user retrieved successfully",
		Data: map[string]any{
			"user": *user.ToResponse(),
		},
	})
}

// GetPublicUser godoc
// @Summary Get public user
// @Description Retrieve the public details of a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} util.SuccessResponse "User retrieved successfully"
// @Failure 400 {object} util.ErrorResponse "Invalid user ID"
// @Failure 500 {object} util.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (u *UserHandler) GetPublicUser(ctx *gin.Context) {
	userID, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		})
	}
	userId, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		})
	}
	user, err := u.userService.GetCurrentUser(ctx.Request.Context(), userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		},
		)
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "User retrieved successfully",
		Data: map[string]any{
			"user": *user.ToPublicResponse(),
		},
	})

}

// UpdateUser godoc
// @Summary Update user
// @Description Update the details of the currently authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.UpdateUserRequest true "User update request"
// @Success 200 {object} util.SuccessResponse "User updated successfully"
// @Failure 400 {object} util.ErrorResponse "Invalid JSON format"
// @Failure 500 {object} util.ErrorResponse "Internal server error"
// @Router /users/me [put]
func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	var req models.UpdateUserRequest
	id := ctx.GetInt("user_id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	if err := u.userService.UpdateUser(ctx.Request.Context(), id, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{

			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "User updated successfully",
	})

}
