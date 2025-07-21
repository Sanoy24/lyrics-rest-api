package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

type UserRepository interface {
	// Define methods that the UserRepository should implement
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
