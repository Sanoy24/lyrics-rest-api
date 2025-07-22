package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status:  false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.authService.Register(ctx.Request.Context(), &req)
	log.Printf("error: %v", err)
	fmt.Println(response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status:  false,
			Message: "Registration failed",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
