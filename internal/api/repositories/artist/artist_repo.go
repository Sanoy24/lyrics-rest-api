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
func (r *artistRepo) GetAllArtists(ctx context.Context) ([]*models.Artist, error) {
	return nil, nil
}

// - GET /api/v1/artists/:id → Get artist details
func (r *artistRepo) GetArtistByID(ctx context.Context, id int) (*models.Artist, error) {
	return nil, nil
}

// - PUT /api/v1/artists/:id → Update artist details
func (r *artistRepo) UpdateArtist(ctx context.Context, id int, artist *models.UpdateArtistRequest) error {
	return nil
}

// - DELETE /api/v1/artists/:id → Delete artist
func (r *artistRepo) DeleteArtist(ctx context.Context, id int) error {
	return nil
}
