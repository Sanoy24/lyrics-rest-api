package models

import "time"

type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Associations
	Users       []User       `json:"users" gorm:"foreignKey:RoleID"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
