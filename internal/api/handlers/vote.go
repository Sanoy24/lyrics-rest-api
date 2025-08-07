package handlers

import (
	"net/http"

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

	v.voteService.CastVote(ctx.Request.Context(), &req, req.Id)

}
