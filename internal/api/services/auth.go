package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	userRepo  interfaces.UserRepository
	jwtExpiry time.Duration
	jwtSecret string
	validator *validator.Validate
	logger    *slog.Logger
}

func NewAuthService(userRepo interfaces.UserRepository, jwtExpiry time.Duration, jwtSecret string, logger *slog.Logger) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtExpiry: jwtExpiry,
		jwtSecret: jwtSecret,
		validator: validator.New(),
		logger:    logger,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.CreateUserRequest) (*util.SuccessResponse, error) {
	if err := s.validateCreateUserRequest(req); err != nil {
		s.logger.Warn("Invalid user registration request",
			slog.String("error", err.Error()),
			slog.String("username", req.Username),
		)
		return nil, err
	}
	if err := s.checkUserExists(ctx, req.Email, req.Username); err != nil {
		return nil, err
	}

	role, err := s.userRepo.GetRoleByName(ctx, "user")
	if err != nil {
		s.logger.Error("Failed to get role", slog.String("error", err.Error()))
		return nil, customerror.ErrInternalServer
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("Failed to hash password", slog.String("error", err.Error()))
		return nil, customerror.ErrInternalServer
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
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error("Failed to create User", slog.String("error", err.Error()), slog.String("username", req.Username), slog.String("email", req.Email))
		if errors.Is(err, customerror.ErrUserExists) {
			return nil, err
		}
		return nil, customerror.ErrInternalServer
	}
	s.logger.Info("User registered successfully", slog.Int("user_id", int(user.ID)), slog.String("username", user.Username))
	return &util.SuccessResponse{
		Status:  true,
		Message: "User created successfully",
	}, nil

}

func (s *AuthService) validateCreateUserRequest(req *models.CreateUserRequest) error {
	if err := s.validator.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				return customerror.NewValidationError(validationErr.Field(), getValidationErrorMessage(validationErr))
			}
		}
		return customerror.ErrInvalidInput
	}
	return nil
}

func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		return fmt.Sprintf("Must be at least %s characters long", err.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s characters long", err.Param())
	case "url":
		return "Must be a valid URL"
	default:
		return "Invalid value"
	}
}

func (s *AuthService) checkUserExists(ctx context.Context, email, username string) error {
	if _, err := s.userRepo.GetUserByEmail(ctx, email); err == nil {
		return customerror.NewAppError("EMAIL_EXISTS", "User with this email already exists", 409)
	} else if !errors.Is(err, customerror.ErrUserNotFound) {
		s.logger.Error("Error checking email existence", slog.String("error", err.Error()))
		return customerror.ErrInternalServer
	}

	if _, err := s.userRepo.GetUserByUsername(ctx, username); err == nil {
		return customerror.NewAppError("USERNAME_EXISTS", "User with this username already exists", 409)
	} else if !errors.Is(err, customerror.ErrUserNotFound) {
		s.logger.Error("Error checking username existence", slog.String("error", err.Error()))
		return customerror.ErrInternalServer
	}

	return nil
}
