package handlers

import (
	"net/http"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadService *services.UploadService
}

func NewUploadHandler(uploadService *services.UploadService) *UploadHandler {
	return &UploadHandler{uploadService: uploadService}
}

func (u *UploadHandler) UploadImageHandler(ctx *gin.Context) {
	imgUrl, err := u.uploadService.UploadImage(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &util.ErrorResponse{
			Status: false,
			Error: util.ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "file upload failed",
				Details: err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &util.SuccessResponse{
		Status:  true,
		Message: "file uploaded successfully",
		Data: map[string]string{
			"url": imgUrl,
		},
	})
}
