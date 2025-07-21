package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Bio       string         `json:"bio"`
	Avatar    string         `json:"avatar"`
	Verified  bool           `json:"verified" gorm:"default:false"`
	Active    bool           `json:"active" gorm:"default:true"`
	RoleID    uint           `json:"role_id"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Association
	Artists     []Artist     `json:"artists" gorm:"foreignKey:UserID"`
	Songs       []Song       `json:"songs" gorm:"foreignKey:ContributorID"`
	Annotations []Annotation `json:"annotations" gorm:"foreignKey:UserID"`
	Comments    []Comment    `json:"comments" gorm:"foreignKey:UserID"`
	Votes       []Vote       `json:"votes" gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,max=100"`
	FirstName string `json:"first_name" binding:"omitempty,max=50"`
	LastName  string `json:"last_name" binding:"omitempty,max=50"`
	Bio       string `json:"bio" binding:"omitempty,max=500"`
	Avatar    string `json:"avatar" binding:"omitempty,url"`
}
type UpdateUserRequest struct {
	Username  string `json:"username" binding:"omitempty,min=3,max=30"` // Optional
	Email     string `json:"email" binding:"omitempty,email"`
	Password  string `json:"password" binding:"omitempty,min=6,max=100"` // Optional
	FirstName string `json:"first_name" binding:"omitempty,max=50"`      // Optional
	LastName  string `json:"last_name" binding:"omitempty,max=50"`       // Optional
	Bio       string `json:"bio" binding:"omitempty,max=500"`            // Optional
	Avatar    string `json:"avatar" binding:"omitempty,url"`             // Optional
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
	Verified  bool   `json:"verified"`
	Active    bool   `json:"active"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"role"`
}

type PaginatedUserResponse struct {
	Data       []UserResponse `json:"data"`
	Total      int64          `json:"total" example:"100"`
	Page       int            `json:"page" example:"1"`
	Limit      int            `json:"limit" example:"10"`
	TotalPages int            `json:"total_pages" example:"10"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Bio:       u.Bio,
		Avatar:    u.Avatar,
		Verified:  u.Verified,
		Active:    u.Active,
		RoleID:    u.RoleID,
		Role:      u.Role,
	}
}
