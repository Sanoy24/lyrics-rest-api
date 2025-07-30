package models

import (
	"strings"
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
	IsDeleted   bool           `json:"is_deleted" gorm:"default:false"`

	// Associations
	Songs       []Song       `json:"songs" gorm:"many2many:song_artists;"`
	Albums      []Album      `json:"albums" gorm:"foreignKey:ArtistID"`
	Annotations []Annotation `json:"annotations" gorm:"foreignKey:ArtistID"`
}

type CreateArtistRequest struct {
	Name        string `json:"name" binding:"required" validate:"required,min=2,max=80"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	Image       string `json:"image,omitempty" validate:"omitempty,url"`
	HeaderImage string `json:"header_image,omitempty" validate:"omitempty,url"`
	Verified    bool   `json:"verified"`
	UserID      int    `json:"user_id,omitempty"`
}

type UpdateArtistRequest struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	HeaderImage string `json:"header_image,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
}

type ArtistResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Verified bool   `json:"verified"`
	UserID   uint   `json:"user_id"`
}

func (a *Artist) ToResponse() *ArtistResponse {
	return &ArtistResponse{
		ID:       a.ID,
		Name:     a.Name,
		Slug:     a.Slug,
		Verified: a.Verified,
		UserID:   a.UserID,
	}
}

func (a *Artist) BeforeCreate(tx *gorm.DB) (err error) {
	a.Slug = strings.Join(strings.Split(a.Name, " "), "-")
	return
}
