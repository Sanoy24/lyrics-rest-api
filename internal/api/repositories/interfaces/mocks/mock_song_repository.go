package mocks

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockSongRepository is a mock implementation of the SongRepository interface
type MockSongRepository struct {
	mock.Mock
}

// CreateSong mocks the CreateSong method
func (m *MockSongRepository) CreateSong(ctx context.Context, song *models.Song) error {
	args := m.Called(ctx, song)
	return args.Error(0)
}

// GetSongByID mocks the GetSongByID method
func (m *MockSongRepository) GetSongByID(ctx context.Context, id int) (*models.Song, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Song), args.Error(1)
}

// GetSongBySlug mocks the GetSongBySlug method
func (m *MockSongRepository) GetSongBySlug(ctx context.Context, slug string) (*models.Song, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Song), args.Error(1)
}

// UpdateSong mocks the UpdateSong method
func (m *MockSongRepository) UpdateSong(ctx context.Context, id int, song *models.UpdateSongRequest) error {
	args := m.Called(ctx, id, song)
	return args.Error(0)
}

// DeleteSong mocks the DeleteSong method
func (m *MockSongRepository) DeleteSong(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// GetAllSongs mocks the GetAllSongs method
func (m *MockSongRepository) GetAllSongs(ctx context.Context, limit, offset int) ([]models.Song, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Song), args.Error(1)
}

// GetSongsCount mocks the GetSongsCount method
func (m *MockSongRepository) GetSongsCount(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

// UpdateVoteScore mocks the UpdateVoteScore method
func (m *MockSongRepository) UpdateVoteScore(ctx context.Context, songID uint, scoreDelta int) error {
	args := m.Called(ctx, songID, scoreDelta)
	return args.Error(0)
}

// SearchSongs mocks the SearchSongs method
func (m *MockSongRepository) SearchSongs(ctx context.Context, query string) ([]models.Song, error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Song), args.Error(1)
}