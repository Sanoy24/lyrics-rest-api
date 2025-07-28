package services

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
)

type UserService struct {
	userRepo interfaces.UserRepository
	logger   *zap.Logger
}

func NewUserService(userRepo interfaces.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *UserService) GetCurrentUser(ctx context.Context, id int) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get user by id", zap.Error(err), zap.Int("user_id", id))
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) error {
	if err := s.userRepo.UpdateUser(ctx, id, req); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}
