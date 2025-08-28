package testutils

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"time"
)

// CreateTestArtist creates a fully initialized artist model for testing
func CreateTestArtist() *models.Artist {
	return &models.Artist{
		ID:          1,
		Name:        "Test Artist",
		Slug:        "test-artist",
		Description: "This is a test artist description",
		Image:       "test-image.jpg",
		HeaderImage: "test-header-image.jpg",
		Verified:    true,
		UserID:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// CreateTestArtists creates a slice of test artists for testing
func CreateTestArtists(count int) []models.Artist {
	artists := make([]models.Artist, count)
	for i := 0; i < count; i++ {
		artists[i] = *CreateTestArtist()
		artists[i].ID = uint(i + 1)
		artists[i].Name = "Test Artist " + string(rune('A'+i))
		artists[i].Slug = "test-artist-" + string(rune('a'+i))
	}
	return artists
}