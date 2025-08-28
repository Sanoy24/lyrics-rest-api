package util

import (
	"context"
	"fmt"
	"os"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UplodaToClaudinary(ctx context.Context, filePath, filenane string, cfg *config.Config) (string, error) {
	defer func() {
		os.Remove(filePath)
	}()
	cloudinary_url := fmt.Sprintf("cloudinary://%s:%s@%s",
		cfg.ClaudinaryKeys.ClaudinaryApiKey,
		cfg.ClaudinaryKeys.ClaudinarySecret,
		cfg.ClaudinaryKeys.ClaudinaryCloudName,
	)
	fmt.Println("claudinary url:", cloudinary_url)
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	if err != nil {
		return "", err
	}
	cld.Config.URL.Secure = true

	res, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		PublicID: filenane,
	})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil

}
