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
	var song models.Song
	if err := s.db.WithContext(ctx).Where("id=?", id).First(&song).Error; err != nil {
		return nil, err
	}
	s.logger.Info("Song fetched successfully", zap.Any("song", id))
	return &song, nil
}
func (s *songRepo) GetSongBySlug(ctx context.Context, slug string) (*models.Song, error) {
	var song models.Song
	if err := s.db.WithContext(ctx).Where("slug=?", slug).First(&song).Error; err != nil {
		s.logger.Error("failed to get song by slug", zap.Error(err))
		return nil, err
	}
	s.logger.Info("Song fetched successfully", zap.Any("song", song))
	return &song, nil
}
func (s *songRepo) UpdateSong(ctx context.Context, id int, song *models.UpdateSongRequest) error {
	if err := s.db.WithContext(ctx).Model(&models.Song{}).Where("id=?", id).Updates(song).Error; err != nil {
		return err
	}
	s.logger.Info("Song updated successfully", zap.Any("song", song))
	return nil
}
func (s *songRepo) DeleteSong(ctx context.Context, id int) error {
	song := &models.Song{ID: uint(id)}
	if err := s.db.WithContext(ctx).Delete(song).Error; err != nil {
		s.logger.Error("failed to delete song", zap.Error(err))
		return err
	}
	s.logger.Info("Song deleted successfully", zap.Int("song_id", id))
	return nil

}
func (s *songRepo) GetAllSongs(ctx context.Context, limit, offset int) ([]models.Song, error) {
	var song []models.Song
	if err := s.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&song).Error; err != nil {
		s.logger.Error("failed to get songs", zap.Error(err))
		return nil, err
	}
	s.logger.Info("Songs fetched successfully", zap.Any("songs", song))
	return song, nil

}

func (r *songRepo) GetSongsCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Song{}).Count(&count).Error; err != nil {
		r.logger.Error("failed to get artists count", zap.Error(err))
		return count, err
	}
	return count, nil
}
