package models

import (
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"unique;not null"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	HeaderImage string         `json:"header_image"`
	Verified    bool           `json:"verified" gorm:"default:false"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Associations
	Songs       []Song       `json:"songs" gorm:"many2many:song_artists;"`
	Albums      []Album      `json:"albums" gorm:"foreignKey:ArtistID"`
	Annotations []Annotation `json:"annotations" gorm:"foreignKey:ArtistID"`
}
