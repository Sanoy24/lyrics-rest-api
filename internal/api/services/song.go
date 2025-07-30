package services

import (
	"context"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
)

type SongService struct {
	songRepo    interfaces.SongRepository
	artistsRepo interfaces.ArtistRepository
	logger      *zap.Logger
}

func NewSongService(songRepo interfaces.SongRepository, artistRepo interfaces.ArtistRepository, logger *zap.Logger) *SongService {
	return &SongService{
		songRepo:    songRepo,
		artistsRepo: artistRepo,
		logger:      logger,
	}
}

func (s *SongService) CreateSong(ctx context.Context, req *models.CreateSongRequest) (*models.Song, error) {

	if req.ReleaseDate == "" {
		req.ReleaseDate = time.Now().Format("2006-01-02")
	}

	song := &models.Song{
		Title:         req.Title,
		Lyrics:        &req.Lyrics,
		Description:   req.Description,
		Image:         req.Image,
		AlbumID:       req.AlbumID,
		ContributorID: req.ContributorID,
		Verified:      req.Verified,
	}

	if len(req.ArtistIDs) > 0 {
		artists, err := s.artistsRepo.GetArtistByIds(ctx, req.ArtistIDs)
		if err != nil {
			s.logger.Error("failed to fetch artists", zap.Error(err))
			return nil, err
		}
		song.Artists = artists
	}

	if err := s.songRepo.CreateSong(ctx, song); err != nil {
		return nil, err

	}
	return song, nil
}
