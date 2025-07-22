package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) interfaces.UserRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
func (r *postgresRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}
func (r *postgresRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Model(user).Updates(user).Error
}
func (r *postgresRepository) DeleteUser(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
func (r *postgresRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *postgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *postgresRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *postgresRepository) GetRoleByName(ctx context.Context, roleName string) (*models.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).Where("name = ?", roleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get role by name: %w", err)
	}
	return &role, nil
}
