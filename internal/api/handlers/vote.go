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

type VoteHandler struct {
	voteService *services.VoteService
	logger      *zap.Logger
}

func NewVoteHandler(voteService *services.VoteService, logger *zap.Logger) *VoteHandler {
	return &VoteHandler{
		voteService: voteService,
		logger:      logger,
	}
}

func (v *VoteHandler) CastVote(ctx *gin.Context) {
	var req models.CreateVoteRequest
	req.UserID = uint(ctx.GetInt("user_id"))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.logger.Error("failed to bind json", zap.Error(err))
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
		return

	}

	if err := v.voteService.CastVote(ctx.Request.Context(), &req, req.Id); err != nil {
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
		Message: "vote created successfully",
	})

}

func (v *VoteHandler) UpdateVoteScore(ctx *gin.Context) {
	entityID, _ := strconv.Atoi(ctx.Param("id"))

	var req models.UpdateVote
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INVALID_REQUEST",
				Message: "invalid request data",
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

	if err := v.voteService.UpdateEntityVoteScore(ctx.Request.Context(), req.EntityType, uint(entityID), int(req.VoteDelta)); err != nil {
		ctx.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "error in updating vote",
				Details: err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "vote updated successfully",
	})
}
