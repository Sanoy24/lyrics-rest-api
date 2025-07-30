package song

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type songRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewSongRepo(db *gorm.DB, logger *zap.Logger) interfaces.SongRepository {
	return &songRepo{
		db:     db,
		logger: logger,
	}
}

func (s *songRepo) CreateSong(ctx context.Context, song *models.Song) error {
	if err := s.db.WithContext(ctx).Create(&song).Error; err != nil {
		return err
	}
	s.logger.Info("Song created successfully", zap.Any("song", song))
	return nil
}
func (s *songRepo) GetSongByID(ctx context.Context, id int) (*models.Song, error) {
	return nil, nil
}
func (s *songRepo) GetSongBySlug(ctx context.Context, slug string) (*models.Song, error) {
	return nil, nil
}
func (s *songRepo) UpdateSong(ctx context.Context, id int, song *models.UpdateSongRequest) error {
	return nil
}
func (s *songRepo) DeleteSong(ctx context.Context, id int) error {
	return nil
}
func (s *songRepo) GetAllSongs(ctx context.Context, limit, offset int) ([]models.Song, error) {
	return nil, nil

}
