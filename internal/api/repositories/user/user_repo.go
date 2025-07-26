package user

import (
	"context"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepo(db *gorm.DB, logger *zap.Logger) interfaces.UserRepository {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		u.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	u.logger.Info("user created successfully", zap.Int("user_id", int(user.ID)), zap.String("username", user.Username))
	return nil

}
func (u *userRepo) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		u.logger.Error("failed to get user by id", zap.Error(err), zap.Int("user_id", int(id)))
		return nil, err
	}
	u.logger.Info("user retrieved successfully", zap.Int("user_id", int(id)), zap.String("username", user.Username))
	return &user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}
func (u *userRepo) DeleteUser(ctx context.Context, id int) error {
	return nil
}
func (u *userRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		u.logger.Error("failed to get user by email", zap.Error(err), zap.String("email", email))
		return nil, err
	}
	u.logger.Info("user retrieved successfully", zap.String("email", email), zap.String("username", user.Username))
	return &user, nil
}
func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}
func (u *userRepo) GetRoleByName(ctx context.Context, roleName string) (*models.Role, error) {
	var role models.Role
	if err := u.db.WithContext(ctx).Where("name=?", roleName).First(&role).Error; err != nil {
		u.logger.Error("failed to get role by name", zap.Error(err), zap.String("role_name", roleName))
		return nil, err
	}
	u.logger.Info("role retrieved successfully", zap.String("role_name", roleName))
	return &role, nil
}

func (u *userRepo) GetByUsernameOrPassword(ctx context.Context, identifier string) (*models.User, error) {
	var user models.User
	if err := u.db.WithContext(ctx).Where("username = ? OR email = ?", identifier, identifier).Preload("Role").First(&user).Error; err != nil {
		u.logger.Error("failed to get user by username or email", zap.Error(err), zap.String("identifier", identifier))
		return nil, err
	}
	u.logger.Info("user retrieved successfully", zap.String("identifier", identifier), zap.String("username", user.Username))
	return &user, nil
}
