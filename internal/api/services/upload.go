package services

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

type UploadService struct {
	cfg *config.Config
}

func NewUploadService(cfg *config.Config) *UploadService {
	return &UploadService{
		cfg: cfg,
	}
}

const (
	maxFileSize = 5 << 20
	uploadDir   = "./assets/uploads/"
)

var allowedMimeTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/gif":  {},
}

func (u UploadService) UploadImage(ctx *gin.Context) (string, error) {
	if err := ctx.Request.ParseMultipartForm(maxFileSize); err != nil {
		return "", err
	}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		return "", err
	}
	fileHeader, err := file.Open()

	if err != nil {

		return "", err
	}
	defer fileHeader.Close()

	head := make([]byte, 400)
	fileHeader.Read(head)
	kind, _ := filetype.Match(head)

	if kind == filetype.Unknown {

		return "", err
	}
	contentType := kind.MIME.Value
	if _, ok := allowedMimeTypes[contentType]; !ok {

		return "", fmt.Errorf("unsupported file type: %s", contentType)
	}

	fileExt := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
	filePath := filepath.Join(uploadDir, newFilename)

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {

		return "", err
	}
	c := context.Background()

	imgUrl, err := util.UplodaToClaudinary(c, filePath, newFilename, u.cfg)
	if err != nil {

		return "", err
	}

	return imgUrl, nil

}
