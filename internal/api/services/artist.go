package services

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"go.uber.org/zap"
)

type ArtistService struct {
	artistRepo interfaces.ArtistRepository
	logger     *zap.Logger
}

func NewArtistService(artistRepo interfaces.ArtistRepository, logger *zap.Logger) *ArtistService {
	return &ArtistService{
		artistRepo: artistRepo,
		logger:     logger,
	}
}

func (a *ArtistService) CreateArtist(ctx context.Context, req *models.CreateArtistRequest, id int) (*util.SuccessResponse, any) {

	if appErr := util.ValidateStruct(req); appErr != nil {
		a.logger.Warn("invalid artist request", zap.Any("error", appErr))
		return nil, appErr
	}

	artist := &models.Artist{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Image:       req.Image,
		HeaderImage: req.HeaderImage,
		Verified:    req.Verified,
		UserID:      uint(id),
	}
	if err := a.artistRepo.CreateArtist(ctx, artist); err != nil {
		a.logger.Error("failed to create artist", zap.Error(err))
		return nil, err
	}
	a.logger.Info("artist created successfully", zap.Int("artist_id", int(artist.ID)), zap.String("name", artist.Name))
	return &util.SuccessResponse{
		Status:  true,
		Message: "artist created successfully",
		Data: map[string]any{
			"artist": *artist.ToResponse(),
		}}, nil
}
