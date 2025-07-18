package models

type ArtistTag struct {
	ArtistID uint `json:"artist_id" gorm:"primaryKey"`
	TagID    uint `json:"tag_id" gorm:"primaryKey"`
}
