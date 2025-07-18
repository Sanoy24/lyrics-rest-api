package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Slug          string         `json:"slug" gorm:"unique;not null"`
	Lyrics        string         `json:"lyrics"`
	Description   string         `json:"description"`
	Image         string         `json:"image"`
	ReleaseDate   *time.Time     `json:"release_date"`
	AlbumID       *uint          `json:"album_id"`
	Album         *Album         `json:"album" gorm:"foreignKey:AlbumID"`
	ContributorID uint           `json:"contributor_id"`
	Contributor   User           `json:"contributor" gorm:"foreignKey:ContributorID"`
	ViewCount     int64          `json:"view_count" gorm:"default:0"`
	LikeCount     int64          `json:"like_count" gorm:"default:0"`
	Verified      bool           `json:"verified" gorm:"default:false"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	// Association
	Artists     []Artist     `json:"artists" gorm:"many2many:song_artists"`
	Annotations []Annotation `json:"annotations" gorm:"foreignKey:SongID"`
	Comments    []Comment    `json:"comments" gorm:"foreignKey:SongID"`
	Votes       []Vote       `json:"votes" gorm:"foreignKey:SongID"`
}
