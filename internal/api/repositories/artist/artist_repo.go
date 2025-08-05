package artist

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type artistRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewArtistRepo(db *gorm.DB, logger *zap.Logger) interfaces.ArtistRepository {
	return &artistRepo{
		db:     db,
		logger: logger,
	}
}

func (r *artistRepo) CreateArtist(ctx context.Context, artist *models.Artist) error {
	if err := r.db.WithContext(ctx).Create(artist).Error; err != nil {
		r.logger.Error("failed to create artist", zap.Error(err))
		return err
	}
	r.logger.Info("artist created successfully", zap.Int("artist_id", int(artist.ID)), zap.String("name", artist.Name))
	return nil
}

// - GET /api/v1/artists → List all artists
func (r *artistRepo) GetAllArtists(ctx context.Context, limit, offset int) ([]models.Artist, error) {
	var artists []models.Artist
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&artists).Error; err != nil {
		r.logger.Error("failed to get artists", zap.Error(err))
		return artists, err
	}
	return artists, nil
}

// - GET /api/v1/artists/:id → Get artist details
func (r *artistRepo) GetArtistByID(ctx context.Context, id int) (*models.Artist, error) {
	var artist models.Artist
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&artist).Error; err != nil {
		r.logger.Error("failed to get artist", zap.Error(err))
		return nil, err
	}
	return &artist, nil
}

// - PUT /api/v1/artists/:id → Update artist details
func (r *artistRepo) UpdateArtist(ctx context.Context, id int, artist *models.UpdateArtistRequest) error {
	if err := r.db.WithContext(ctx).Model(&models.Artist{}).Where("id = ?", id).Updates(artist).Error; err != nil {
		r.logger.Error("failed to update artist", zap.Error(err))
		return err
	}
	return nil
}

// - DELETE /api/v1/artists/:id → Delete artist
func (r *artistRepo) DeleteArtist(ctx context.Context, id int) error {
	artist := &models.Artist{ID: uint(id)}
	if err := r.db.WithContext(ctx).Delete(artist).Error; err != nil {
		r.logger.Error("failed to delete artist", zap.Error(err))
		return err
	}
	r.logger.Info("Artist Deleted Successfully")
	return nil
}

func (r *artistRepo) GetArtistsCount(ctx context.Context) (int64, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(&models.Artist{}).Count(&count).Error; err != nil {
		r.logger.Error("failed to get artists count", zap.Error(err))
		return count, err
	}
	return count, nil

}

func (r *artistRepo) GetArtistByIds(ctx context.Context, ids []int) ([]models.Artist, error) {
	var artists []models.Artist
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&artists).Error; err != nil {
		r.logger.Error("failed to get artists by ids", zap.Error(err))
		return artists, err
	}
	r.logger.Info("artists fetched successfully", zap.Any("artists", artists))
	return artists, nil
}
