package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Content      string         `json:"content" gorm:"type:text;not null"`
	UserID       uint           `json:"user_id"`
	User         User           `json:"user" gorm:"foreignKey:UserID"`
	SongID       *uint          `json:"song_id"`
	Song         *Song          `json:"song" gorm:"foreignKey:SongID"`
	AnnotationID *uint          `json:"annotation_id"`
	Annotation   *Annotation    `json:"annotation" gorm:"foreignKey:AnnotationID"`
	ParentID     *uint          `json:"parent_id"`
	Parent       *Comment       `json:"parent" gorm:"foreignKey:ParentID"`
	VoteScore    int            `json:"vote_score" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Annotations
	Replies []Comment `json:"replies" gorm:"foreignKey:ParentID"`
	Votes   []Vote    `json:"votes" gorm:"foreignKey:CommentID"`
}
