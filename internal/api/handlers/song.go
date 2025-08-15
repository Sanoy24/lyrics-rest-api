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

type SongHandler struct {
	songService *services.SongService
	logger      *zap.Logger
}

func NewSongHandler(songService *services.SongService, logger *zap.Logger) *SongHandler {
	return &SongHandler{
		songService: songService,
		logger:      logger,
	}
}

func (s *SongHandler) CreateSong(ctx *gin.Context) {
	var req models.CreateSongRequest
	id := ctx.Value("user_id").(int)
	req.ContributorID = uint(id)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.logger.Error("failed to bind json", zap.Error(err))
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
	}
	response, err := s.songService.CreateSong(ctx.Request.Context(), &req)
	if err != nil {

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
	ctx.JSON(http.StatusCreated, util.SuccessResponse{
		Status: true,
		Data:   response,
	})
}

func (s *SongHandler) GetAllSongs(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	data, err := s.songService.GetAllSongs(ctx.Request.Context(), limit, offset)

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
	ctx.JSON(http.StatusOK, data)
}

func (s *SongHandler) GetSongById(ctx *gin.Context) {
	id := ctx.Param("id")
	songID, _ := strconv.Atoi(id)
	response, err := s.songService.GetSongById(ctx.Request.Context(), songID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "NOT_FOUND",
				Message: "Song not found",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "song data fetched successfully",
		Data:    response,
	})
}

func (s *SongHandler) UpdateSong(ctx *gin.Context) {
	var req models.UpdateSongRequest
	id := ctx.Param("id")
	songID, _ := strconv.Atoi(id)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INVALID_JSON",
				Message: "Invalid payload",
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
	if err := s.songService.UpdateSong(ctx.Request.Context(), songID, &req); err != nil {
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
		Message: "song updated successfully",
	})

}

func (s *SongHandler) GetSongBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, err := s.songService.GetSongBySlug(ctx.Request.Context(), slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
				Details: err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse{
		Status:  true,
		Message: "song fetched successfully",
		Data:    response,
	})
}

func (s *SongHandler) DeleteSong(ctx *gin.Context) {
	id := ctx.Param("id")
	songID, _ := strconv.Atoi(id)
	if err := s.songService.DeleteSong(ctx.Request.Context(), songID); err != nil {
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
		Message: "song deleted successfully",
	})

}

func (s *SongHandler) SearchSongs(ctx *gin.Context) {
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "BAD_REQUEST",
				Message: "enter a valid search query ",
			},
		})
		return
	}
	songs, err := s.songService.SearchSong(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
			Status: false,
			Error: &util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "error while doing search",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &util.SuccessResponse{
		Status:  true,
		Message: "song fetched by query successfully",
		Data:    songs,
	})

}
