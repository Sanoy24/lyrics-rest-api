package services

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

type UserService struct {
	userRepo  interfaces.UserRepository
	jwtExpiry string
	jwtSecret string
}

func NewUserService(userRepo interfaces.UserRepository, jwtExpiry, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtExpiry: jwtExpiry,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) Register(ctx context.Context, req *models.CreateUserRequest) interfaces.UserRepository {

}
