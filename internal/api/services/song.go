package services

import (
	"context"
	"math"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
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
		Lyrics:        req.Lyrics,
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

func (s *SongService) GetSongById(ctx context.Context, id int) (*models.Song, error) {
	song, err := s.songRepo.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (s *SongService) UpdateSong(ctx context.Context, id int, req *models.UpdateSongRequest) error {
	if err := s.songRepo.UpdateSong(ctx, id, req); err != nil {
		return err
	}
	return nil
}

func (s *SongService) GetAllSongs(ctx context.Context, limit, offset int) (*util.PaginatedResponse, error) {

	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}
	totalCount, err := s.songRepo.GetSongsCount(ctx)

	if err != nil {
		s.logger.Info("failed to get songs count", zap.Error(err))
		return nil, err
	}

	songs, err := s.songRepo.GetAllSongs(ctx, limit, offset)
	if err != nil {
		s.logger.Info("failed to get songs", zap.Error(err))
		return nil, err
	}

	songResponse := make([]models.SongResponse, 0, len(songs))
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	currentPage := (offset / limit) + 1
	hasNext := currentPage < totalPages
	hasPrev := currentPage > 1

	for _, songs := range songs {
		songResponse = append(songResponse, *songs.ToResponse())
	}

	s.logger.Info("songs fetched successfully")
	return &util.PaginatedResponse{
		Sucess:  true,
		Message: "songs fetched successfully",
		Data:    songResponse,
		Pagination: &util.Pagination{
			Page:        int(totalPages),
			Limit:       int(limit),
			Total:       int(totalCount),
			TotalPages:  totalPages,
			HasNext:     hasNext,
			HasPrev:     hasPrev,
			CurrentPage: currentPage,
		},
	}, nil
}

func (s *SongService) GetSongBySlug(ctx context.Context, slug string) (*models.Song, error) {
	song, err := s.songRepo.GetSongBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return song, nil
}
