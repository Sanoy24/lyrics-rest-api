package services

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/errors"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
)

type AuthService struct {
	userRepo  interfaces.UserRepository
	jwtExpiry string
	jwtSecret string
}

func NewAuthService(userRepo interfaces.UserRepository, jwtExpiry, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtExpiry: jwtExpiry,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.CreateUserRequest) (*util.SuccessResponse, error) {
	if _, err := s.userRepo.GetUserByEmail(ctx, req.Email); err == nil {
		return nil, errors.ErrUserExists
	}
	if _, err := s.userRepo.GetUserByUsername(ctx, req.Username); err == nil {
		return nil, errors.ErrUserExists
	}
	// check if role exists

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	role, err := s.userRepo.GetRoleByName(ctx, "user")
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		RoleID:    role.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Avatar:    req.Avatar,
	}
	// Username  string `json:"username" binding:"required,min=3,max=30"`
	// Email     string `json:"email" binding:"required,email"`
	// Password  string `json:"password" binding:"required,min=6,max=100"`
	// FirstName string `json:"first_name" binding:"omitempty,max=50"`
	// LastName  string `json:"last_name" binding:"omitempty,max=50"`
	// Avatar    string `json:"avatar" binding:"omitempty,url"`
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	return &util.SuccessResponse{
		Status:  true,
		Message: "User created successfully",
	}, nil

}
