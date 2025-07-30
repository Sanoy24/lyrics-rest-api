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

func (a *ArtistService) GetAllArtists(ctx context.Context) (*util.SuccessResponse, error) {
	artists, err := a.artistRepo.GetAllArtists(ctx)
	if err != nil {
		a.logger.Error("failed to get artists", zap.Error(err))
		return nil, err
	}

	var artistResponse []models.ArtistResponse

	for _, artist := range artists {
		artistResponse = append(artistResponse, *artist.ToResponse())

	}
	a.logger.Info("artists fetched successfully")
	return &util.SuccessResponse{
		Status:  true,
		Message: "artists fetched successfully",
		Data: map[string]any{
			"artists": artistResponse,
		}}, nil

}

func (a *ArtistService) GetArtistByID(ctx context.Context, id int) (*util.SuccessResponse, error) {
	artist, err := a.artistRepo.GetArtistByID(ctx, id)
	if err != nil {
		a.logger.Error("failed to get artist", zap.Error(err))
		return nil, err
	}
	a.logger.Info("artist fetched successfully", zap.Int("artist_id", id))
	return &util.SuccessResponse{
		Status:  true,
		Message: "artist fetched successfully",
		Data: map[string]any{
			"artist": *artist.ToResponse(),
		}}, nil
}

func (a *ArtistService) UpdateArtist(ctx context.Context, id int, updateData *models.UpdateArtistRequest) (*util.SuccessResponse, any) {
	if appErr := util.ValidateStruct(updateData); appErr != nil {
		a.logger.Warn("invalid artist request", zap.Any("error", appErr))
		return nil, appErr
	}
	if err := a.artistRepo.UpdateArtist(ctx, id, updateData); err != nil {
		a.logger.Info("failed to update artist", zap.Error(err))
		return nil, err
	}
	return &util.SuccessResponse{
		Status:  true,
		Message: "artist updated successfully",
	}, nil

}

func (a *ArtistService) DeleteArtist(ctx context.Context, id int) (*util.SuccessResponse, error) {
	if err := a.artistRepo.DeleteArtist(ctx, id); err != nil {
		a.logger.Error("failed to delete artist", zap.Error(err))
		return nil, err
	}
	a.logger.Info("artist deleted successfully", zap.Int("artist_id", id))
	return &util.SuccessResponse{
		Status:  true,
		Message: "artist deleted successfully",
	}, nil
}
