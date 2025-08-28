package mocks

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockArtistRepository is a mock implementation of the ArtistRepository interface
type MockArtistRepository struct {
	mock.Mock
}

// CreateArtist mocks the CreateArtist method
func (m *MockArtistRepository) CreateArtist(ctx context.Context, artist *models.Artist) error {
	args := m.Called(ctx, artist)
	return args.Error(0)
}

// GetAllArtists mocks the GetAllArtists method
func (m *MockArtistRepository) GetAllArtists(ctx context.Context, limit, offset int) ([]models.Artist, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Artist), args.Error(1)
}

// GetArtistByID mocks the GetArtistByID method
func (m *MockArtistRepository) GetArtistByID(ctx context.Context, id int) (*models.Artist, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Artist), args.Error(1)
}

// UpdateArtist mocks the UpdateArtist method
func (m *MockArtistRepository) UpdateArtist(ctx context.Context, id int, artist *models.UpdateArtistRequest) error {
	args := m.Called(ctx, id, artist)
	return args.Error(0)
}

// DeleteArtist mocks the DeleteArtist method
func (m *MockArtistRepository) DeleteArtist(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// GetArtistsCount mocks the GetArtistsCount method
func (m *MockArtistRepository) GetArtistsCount(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

// GetArtistByIds mocks the GetArtistByIds method
func (m *MockArtistRepository) GetArtistByIds(ctx context.Context, ids []int) ([]models.Artist, error) {
	args := m.Called(ctx, ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Artist), args.Error(1)
}