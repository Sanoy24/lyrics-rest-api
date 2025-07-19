package data

import "github.com/Sanoy24/lyrics-rest-api/internal/models"

var DefaultRoles = []models.Role{
	{Name: "admin", Description: "Full system access"},
	{Name: "moderator", Description: "Can moderate content and users"},
	{Name: "verified_user", Description: "Verified user with enhanced privileges"},
	{Name: "user", Description: "Regular user"},
	{Name: "guest", Description: "Read-only access"},
}
