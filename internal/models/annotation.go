package models

import (
	"time"

	"gorm.io/gorm"
)

type Annotation struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	SongID     uint           `json:"song_id"`
	Song       Song           `json:"song" gorm:"foreignKey:SongID"`
	ArtistID   *uint          `json:"artist_id"`
	Artist     *Artist        `json:"artist" gorm:"foreignKey:ArtistID"`
	UserID     uint           `json:"user_id"`
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	StartIndex int            `json:"start_index"`
	EndIndex   int            `json:"end_index"`
	Fragment   string         `json:"fragment"`
	Content    string         `json:"content" gorm:"type:text"`
	Verified   bool           `json:"verified" gorm:"default:false"`
	VoteScore  int            `json:"vote_score" gorm:"default:0"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	// Association
	Comments []Comment `json:"comments" gorm:"foreignKey:AnnotationID"`
	Votes    []Vote    `json:"votes" gorm:"foreignKey:AnnotationID"`
}

type CreateAnnotationRequest struct {
	StartIndex int    `json:"start_index" binding:"required"`
	EndIndex   int    `json:"end_index" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type UpdateAnnotationRequest struct {
	StartIndex *int   `json:"start_index,omitempty"`
	EndIndex   *int   `json:"end_index,omitempty"`
	Content    string `json:"content" binding:"required"`
}

type AnnotationSummary struct {
	Count      int
	AvgVotes   float64
	Verified   int
	Unverified int
}
