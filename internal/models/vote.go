package models

import "time"

type Vote struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	UserID       uint        `json:"user_id"`
	User         User        `json:"user" gorm:"foreignKey:UserID"`
	VoteType     string      `json:"vote_type"`
	SongID       *uint       `json:"song_id"`
	Song         *Song       `json:"song" gorm:"foreignKey:SongID"`
	AnnotationID *uint       `json:"annotation_id"`
	Annotation   *Annotation `json:"annotation" gorm:"foreignKey:AnnotationID"`
	CommentID    *uint       `json:"comment_id"`
	Comment      *Comment    `json:"comment" gorm:"foreignKey:CommentID"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
