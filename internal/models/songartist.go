package models

// SongArtist represents the many-to-many relationship between songs and artists
type SongArtist struct {
	SongID   uint   `json:"song_id" gorm:"primaryKey"`
	ArtistID uint   `json:"artist_id" gorm:"primaryKey"`
	Role     string `json:"role"`
}
