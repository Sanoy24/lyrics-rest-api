package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"unique;not null"`
	Description string         `json:"description"`
	CoverArt    string         `json:"cover_art"`
	ReleaseDate *time.Time     `json:"release_date"`
	ArtistID    uint           `json:"artist_id"`
	Artist      Artist         `json:"artist" gorm:"foreignKey:ArtistID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	//Association
	Songs []Song `json:"songs" gorm:"foreignKey:AlbumID"`
}
