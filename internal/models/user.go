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
