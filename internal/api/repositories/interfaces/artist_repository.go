package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

type ArtistRepository interface {
	// Define methods that the SongRepository should implement
	// - POST /api/v1/artists → Add new artist (admin/mod)
	CreateArtist(ctx context.Context, artist *models.Artist) error
	// - GET /api/v1/artists → List all artists
	GetAllArtists(ctx context.Context, limit, offset int) ([]models.Artist, error)
	// - GET /api/v1/artists/:id → Get artist details
	GetArtistByID(ctx context.Context, id int) (*models.Artist, error)
	// - PUT /api/v1/artists/:id → Update artist details
	UpdateArtist(ctx context.Context, id int, artist *models.UpdateArtistRequest) error
	// - DELETE /api/v1/artists/:id → Delete artist
	DeleteArtist(ctx context.Context, id int) error
	// get artists count
	GetArtistsCount(ctx context.Context) (int64, error)
}
