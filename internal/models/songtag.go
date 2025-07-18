package models

// SongTag represents the many-to-many relationship between songs and tags
type SongTag struct {
	SongID uint `json:"song_id" gorm:"primaryKey"`
	TagID  uint `json:"tag_id" gorm:"primaryKey"`
}
