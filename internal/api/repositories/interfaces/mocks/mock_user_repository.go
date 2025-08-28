package mocks

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// CreateUser mocks the CreateUser method
func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// GetUserByID mocks the GetUserByID method
func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// UpdateUser mocks the UpdateUser method
func (m *MockUserRepository) UpdateUser(ctx context.Context, id int, user *models.UpdateUserRequest) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

// DeleteUser mocks the DeleteUser method
func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// GetAllUsers mocks the GetAllUsers method
func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.User), args.Error(1)
}

// GetUserByEmail mocks the GetUserByEmail method
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// GetUserByUsername mocks the GetUserByUsername method
func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// GetRoleByName mocks the GetRoleByName method
func (m *MockUserRepository) GetRoleByName(ctx context.Context, roleName string) (*models.Role, error) {
	args := m.Called(ctx, roleName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Role), args.Error(1)
}

// GetByUsernameOrPassword mocks the GetByUsernameOrPassword method
func (m *MockUserRepository) GetByUsernameOrPassword(ctx context.Context, identifier string) (*models.User, error) {
	args := m.Called(ctx, identifier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}