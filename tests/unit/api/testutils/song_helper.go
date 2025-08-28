package testutils

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"time"
)

// CreateTestSong creates a fully initialized song model for testing
func CreateTestSong() *models.Song {
	releaseDate := time.Now()
	var albumID uint = 1
	return &models.Song{
		ID:            1,
		Title:         "Test Song",
		Slug:          "test-song",
		Lyrics:        "This is a test song lyrics",
		Description:   "Test song description",
		Image:         "test-image.jpg",
		ReleaseDate:   &releaseDate,
		AlbumID:       &albumID,
		ContributorID: 1,
		ViewCount:     100,
		LikeCount:     50,
		Verified:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

// CreateTestSongs creates a slice of test songs for testing
func CreateTestSongs(count int) []models.Song {
	songs := make([]models.Song, count)
	for i := 0; i < count; i++ {
		songs[i] = *CreateTestSong()
		songs[i].ID = uint(i + 1)
		songs[i].Title = "Test Song " + string(rune('A'+i))
		songs[i].Slug = "test-song-" + string(rune('a'+i))
	}
	return songs
}