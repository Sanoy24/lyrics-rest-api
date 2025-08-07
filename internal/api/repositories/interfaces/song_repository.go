package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

/*
	POST /songs - Create songs (essential for content)
	GET /songs/:id - Get single song with lyrics
	PUT /songs/:id - Update song details
	DELETE /songs/:id - Delete songs
	GET /songs - List songs with pagination/filtering
*/

type SongRepository interface {
	// Define methods that the SongRepository should implement
	CreateSong(ctx context.Context, song *models.Song) error
	GetSongByID(ctx context.Context, id int) (*models.Song, error)
	GetSongBySlug(ctx context.Context, slug string) (*models.Song, error)
	UpdateSong(ctx context.Context, id int, song *models.UpdateSongRequest) error
	DeleteSong(ctx context.Context, id int) error
	GetAllSongs(ctx context.Context, limit, offset int) ([]models.Song, error)
	GetSongsCount(ctx context.Context) (int64, error)
	UpdateVoteScore(ctx context.Context, songID uint, scoreDelta int) error
}
