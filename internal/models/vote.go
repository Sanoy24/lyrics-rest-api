package models

import "time"

type Vote struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	UserID       uint        `json:"user_id" gorm:"not null;index:idx_user_entity"`
	User         User        `json:"user" gorm:"foreignKey:UserID"`
	VoteType     string      `json:"vote_type"`
	SongID       *uint       `json:"song_id"`
	Song         *Song       `json:"song" gorm:"foreignKey:SongID"`
	AnnotationID *uint       `json:"annotation_id"`
	Annotation   *Annotation `json:"annotation" gorm:"foreignKey:AnnotationID"`
	CommentID    *uint       `json:"comment_id"`
	Comment      *Comment    `json:"comment" gorm:"foreignKey:CommentID"`
	EntityID     uint        `json:"entity_id" gorm:"not null;index:idx_user_entity"`
	EntityType   string      `json:"entity_type" gorm:"not null;index:idx_user_entity"`
	Value        int8        `json:"value"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateVoteRequest struct {
	UserID     uint   `json:"-"`
	EntityType string `json:"entity_type" binding:"required,oneof=song annotation comment"`
	EntityID   uint   `json:"entity_id" binding:"required"`
	VoteType   string `json:"vote_type" binding:"required,oneof=upvote downvote unvote"`
	Id         int    `json:"-"`
}

type UpdateVote struct {
	EntityType string `json:"entity_type" binding:"required"`
	VoteDelta  int8   `json:"vote_delta" binding:"required"`
}
