package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces/mocks"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/tests/unit/api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestGetSongById(t *testing.T) {
	tests := []struct {
		name      string
		songID    int
		setupMock func(*mocks.MockSongRepository, *mocks.MockArtistRepository)
		wantErr   bool
		errType   error
	}{
		{
			name:   "Success - Song found",
			songID: 1,
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				testSong := testutils.CreateTestSong()
				mockSongRepo.On("GetSongByID", mock.Anything, 1).Return(testSong, nil)
			},
			wantErr: false,
		},
		{
			name:   "Error - Song not found",
			songID: 999,
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				mockSongRepo.On("GetSongByID", mock.Anything, 999).Return(nil, customerror.ErrNotFound)
			},
			wantErr: true,
			errType: customerror.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockSongRepo := new(mocks.MockSongRepository)
			mockArtistRepo := new(mocks.MockArtistRepository)
			tt.setupMock(mockSongRepo, mockArtistRepo)

			// Create logger
			logger, _ := zap.NewDevelopment()

			// Create service
			songService := services.NewSongService(mockSongRepo, mockArtistRepo, logger)

			// Call the method
			song, err := songService.GetSongById(context.Background(), tt.songID)

			// Assert results
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, song)
				assert.Equal(t, uint(tt.songID), song.ID)
			}

			// Verify all expectations were met
			mockSongRepo.AssertExpectations(t)
			mockArtistRepo.AssertExpectations(t)
		})
	}
}

func TestCreateSong(t *testing.T) {
	tests := []struct {
		name      string
		request   *models.CreateSongRequest
		setupMock func(*mocks.MockSongRepository, *mocks.MockArtistRepository)
		wantErr   bool
		errType   error
	}{
		{
			name: "Success - Create song without artists",
			request: &models.CreateSongRequest{
				Title:         "Test Song",
				Lyrics:        "Test lyrics",
				Description:   "Test description",
				ReleaseDate:   time.Now().Format("2006-01-02"),
				ContributorID: 1,
				Verified:      true,
			},
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				mockSongRepo.On("CreateSong", mock.Anything, mock.AnythingOfType("*models.Song")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Success - Create song with artists",
			request: &models.CreateSongRequest{
				Title:         "Test Song with Artists",
				Lyrics:        "Test lyrics",
				Description:   "Test description",
				ReleaseDate:   time.Now().Format("2006-01-02"),
				ContributorID: 1,
				Verified:      true,
				ArtistIDs:     []int{1, 2},
			},
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				artists := testutils.CreateTestArtists(2)
				mockArtistRepo.On("GetArtistByIds", mock.Anything, []int{1, 2}).Return(artists, nil)
				mockSongRepo.On("CreateSong", mock.Anything, mock.AnythingOfType("*models.Song")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error - Failed to fetch artists",
			request: &models.CreateSongRequest{
				Title:         "Test Song with Artists",
				Lyrics:        "Test lyrics",
				Description:   "Test description",
				ReleaseDate:   time.Now().Format("2006-01-02"),
				ContributorID: 1,
				Verified:      true,
				ArtistIDs:     []int{1, 2},
			},
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				mockArtistRepo.On("GetArtistByIds", mock.Anything, []int{1, 2}).Return(nil, customerror.ErrNotFound)
			},
			wantErr: true,
			errType: customerror.ErrNotFound,
		},
		{
			name: "Error - Failed to create song",
			request: &models.CreateSongRequest{
				Title:         "Test Song",
				Lyrics:        "Test lyrics",
				Description:   "Test description",
				ReleaseDate:   time.Now().Format("2006-01-02"),
				ContributorID: 1,
				Verified:      true,
			},
			setupMock: func(mockSongRepo *mocks.MockSongRepository, mockArtistRepo *mocks.MockArtistRepository) {
				mockSongRepo.On("CreateSong", mock.Anything, mock.AnythingOfType("*models.Song")).Return(customerror.ErrInternalServer)
			},
			wantErr: true,
			errType: customerror.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockSongRepo := new(mocks.MockSongRepository)
			mockArtistRepo := new(mocks.MockArtistRepository)
			tt.setupMock(mockSongRepo, mockArtistRepo)

			// Create logger
			logger, _ := zap.NewDevelopment()

			// Create service
			songService := services.NewSongService(mockSongRepo, mockArtistRepo, logger)

			// Call the method
			song, err := songService.CreateSong(context.Background(), tt.request)

			// Assert results
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
				assert.Nil(t, song)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, song)
				assert.Equal(t, tt.request.Title, song.Title)
				assert.Equal(t, tt.request.Lyrics, song.Lyrics)
			}

			// Verify all expectations were met
			mockSongRepo.AssertExpectations(t)
			mockArtistRepo.AssertExpectations(t)
		})
	}
}