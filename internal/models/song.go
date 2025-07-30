package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Slug          string         `json:"slug" gorm:"unique;not null"`
	Lyrics        *string        `json:"lyrics"`
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

type CreateSongRequest struct {
	Title         string `json:"title" binding:"required" validate:"required,min=2,max=80"`
	Lyrics        string `json:"lyrics" binding:"required" validate:"required"`
	Description   string `json:"description,omitempty" validate:"omitempty"`
	Image         string `json:"image,omitempty" validate:"omitempty,url"`
	ReleaseDate   string `json:"release_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	AlbumID       uint   `json:"album_id,omitempty" validate:"omitempty,min=1"`
	ContributorID uint   `json:"contributor_id,omitempty" validate:"omitempty,min=1"`
	Verified      bool   `json:"verified"`
}

type UpdateSongRequest struct {
	Title         string `json:"title,omitempty"`
	Slug          string `json:"slug,omitempty"`
	Lyrics        string `json:"lyrics,omitempty"`
	Description   string `json:"description,omitempty"`
	Image         string `json:"image,omitempty"`
	ReleaseDate   string `json:"release_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	AlbumID       uint   `json:"album_id,omitempty"`
	ContributorID uint   `json:"contributor_id,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
}

func (s *Song) BeforeCreate(tx *gorm.DB) (err error) {
	s.Slug = fmt.Sprintf("%s-%s-%s", strings.Join(strings.Split(s.Artists[0].Name, " "), "-"), strings.Join(strings.Split(s.Title, " "), "-"), "lyrics")
	return
}
