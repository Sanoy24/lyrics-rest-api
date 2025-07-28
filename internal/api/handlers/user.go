package handlers

import (
	"net/http"
	"strconv"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
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

// GET current user /users/me
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

// PUT update user profile /users/me
// GET get public user /users/:id
