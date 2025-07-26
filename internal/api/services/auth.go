package services

import (
	"context"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	customerror "github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"go.uber.org/zap"
)

type AuthService struct {
	userRepo  interfaces.UserRepository
	jwtSecret string
	jwtExpiry time.Duration
	logger    *zap.Logger
}

func NewUserService(userRepo interfaces.UserRepository, jwtSecret string, jwtExpiry time.Duration, logger *zap.Logger) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
		logger:    logger,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.CreateUserRequest) (*util.SuccessResponse, error) {
	if _, err := s.userRepo.GetUserByEmail(ctx, req.Email); err == nil {
		s.logger.Info("user already exists", zap.String("email", req.Email))
		return nil, customerror.ErrUserExists
	}
	hash, err := util.HashPassword(req.Password)
	if err != nil {
		s.logger.Info("failed to hash password", zap.Error(err))
		return nil, customerror.ErrInternalServer
	}
	role, err := s.userRepo.GetRoleByName(ctx, "user")
	s.logger.Info("role", zap.Int("role", int(role.ID)))
	if err != nil {
		s.logger.Info("failed to get role", zap.Error(err))
		return nil, customerror.ErrInternalServer
	}
	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hash,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Avatar:    req.Avatar,
		RoleID:    role.ID,
	}
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, customerror.ErrInternalServer
	}
	return &util.SuccessResponse{
		Status:  true,
		Message: "user created successfully",
		Data: map[string]any{
			"user": *user.ToResponse(),
		},
	}, nil

}

func (s *AuthService) Login(ctx context.Context, req *models.UserLoginRequest) (*util.AuthResponse, error) {
	user, err := s.userRepo.GetByUsernameOrPassword(ctx, req.Identifier)
	if err != nil {
		s.logger.Info("failed to get user by username or email", zap.Error(err))
		return nil, customerror.ErrInvalidCredentials
	}
	if !util.VerifyPassword(user.Password, req.Password) {
		s.logger.Info("invalid password", zap.String("identifier", req.Identifier))
		return nil, customerror.ErrInvalidCredentials
	}
	token, err := util.GenerateJWT(int(user.ID), user.Email, user.Role.Name, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		s.logger.Error("failed to generate jwt token", zap.Error(err))
		return nil, customerror.ErrInternalServer
	}
	s.logger.Info("user logged in successfully", zap.String("indentifier", req.Identifier))
	return &util.AuthResponse{
		Token: token,
		User:  *user.ToResponse(),
	}, nil

}

func (s *AuthService) Logout() {

}

func (s *AuthService) RefreshToken() {

}

func (s *AuthService) ForgotPassword() {

}

func (s *AuthService) ResetPassword() {

}

func (s *AuthService) ChangePassword() {

}

func (s *AuthService) DeleteAccount() {

}
