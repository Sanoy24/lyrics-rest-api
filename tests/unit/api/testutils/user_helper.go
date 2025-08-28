package testutils

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

// CreateTestUser creates a fully initialized user model for testing
func CreateTestUser() *models.User {
	return &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
		FirstName: "Test",
		LastName:  "User",
		Bio:       "Test bio",
		Avatar:    "https://example.com/avatar.jpg",
		Verified:  true,
		Active:    true,
		RoleID:    1,
		Role: models.Role{
			ID:   1,
			Name: "user",
			Permissions: []models.Permission{
				{Resource: "songs", Action: "read"},
			},
		},
	}
}

// CreateTestRole creates a test role for testing
func CreateTestRole() *models.Role {
	return &models.Role{
		ID:   1,
		Name: "user",
		Permissions: []models.Permission{
			{Resource: "songs", Action: "read"},
		},
	}
}