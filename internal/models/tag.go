package models

import "time"

type Tag struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Associations
	Songs   []Song   `json:"songs" gorm:"many2many:song_tags;"`
	Artists []Artist `json:"artists" gorm:"many2many:artist_tags;"`
}
