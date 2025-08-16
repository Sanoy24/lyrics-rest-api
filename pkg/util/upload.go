package util

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

const (
	maxFileSize = 5 << 20
	uploadDir   = "./assets/uploads/"
)

var allowedMimeTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/gif":  {},
}

func UploadImage(ctx *gin.Context, cfg *config.Config) {
	if err := ctx.Request.ParseMultipartForm(maxFileSize); err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "FILE_TOO_LARGE",
				Message: "file exceeds max file size",
			},
		})
		return
	}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "file upload failed",
			},
		})
		return
	}
	fileHeader, err := file.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "failed to open file header",
			},
		})
		return
	}
	defer fileHeader.Close()

	head := make([]byte, 400)
	fileHeader.Read(head)
	kind, _ := filetype.Match(head)

	if kind == filetype.Unknown {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "INVALID_FILE_TYPE",
				Message: "unknown file type",
			},
		})
		return
	}
	contentType := kind.MIME.Value
	if _, ok := allowedMimeTypes[contentType]; !ok {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "INVALID_FILE_TYPE",
				Message: fmt.Sprintf("unsupported file type: %s", contentType),
			},
		})
		return
	}

	fileExt := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
	filePath := filepath.Join(uploadDir, newFilename)

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Error: ErrorData{
				Code:    "INTERNAL_ERROR",
				Message: "file upload failed",
			},
		})
		return
	}
	c := context.Background()

	uplodaToClaudinary()
	// ctx.JSON(http.StatusOK, SuccessResponse{
	// 	Status:  true,
	// 	Message: "file uploaded successfully",
	// })

}

func uplodaToClaudinary(file *multipart.FileHeader, ctx context.Context, cfg *config.Config, assetPath, filename string) (string, error) {
	defer func() {
		os.Remove(assetPath + filename)
	}()
	cloudinary_url := fmt.Sprintf("cloudinary:%s//:%s@%s",
		cfg.ClaudinaryKeys.ClaudinaryApiKey,
		cfg.ClaudinaryKeys.ClaudinarySecret,
		cfg.ClaudinaryKeys.ClaudinaryCloudName,
	)
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	if err != nil {
		return "", err
	}
	cld.Config.URL.Secure = true

	res, err := cld.Upload.Upload(ctx, assetPath+filename, uploader.UploadParams{
		PublicID: filename,
	})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil

}
