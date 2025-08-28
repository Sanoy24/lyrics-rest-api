package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces/mocks"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/tests/unit/api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestAuthService_Register(t *testing.T) {
	// Setup logger
	logger, _ := zap.NewDevelopment()

	// Test cases
	tests := []struct {
		name          string
		setupMock     func(mockRepo *mocks.MockUserRepository)
		input         *models.CreateUserRequest
		expectedError error
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// User doesn't exist yet
				mockRepo.On("GetUserByEmail", mock.Anything, "new@example.com").Return(nil, customerror.ErrUserNotFound)
				
				// Role exists
				mockRole := testutils.CreateTestRole()
				mockRepo.On("GetRoleByName", mock.Anything, "user").Return(mockRole, nil)
				
				// CreateUser succeeds
			mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.User")).Run(func(args mock.Arguments) {
				user := args.Get(1).(*models.User)
				user.ID = 1
				user.Role = *mockRole
			}).Return(nil)
			},
			input: &models.CreateUserRequest{
				Username:  "newuser",
				Email:     "new@example.com",
				Password:  "password123",
				FirstName: "New",
				LastName:  "User",
			},
			expectedError: nil,
		},
		{
			name: "User Already Exists",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// User already exists
				existingUser := testutils.CreateTestUser()
				existingUser.Username = "existinguser"
				existingUser.Email = "existing@example.com"
				mockRepo.On("GetUserByEmail", mock.Anything, "existing@example.com").Return(existingUser, nil)
			},
			input: &models.CreateUserRequest{
				Username:  "existinguser",
				Email:     "existing@example.com",
				Password:  "password123",
				FirstName: "Existing",
				LastName:  "User",
			},
			expectedError: customerror.ErrUserExists,
		},
		{
			name: "Role Not Found",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// User doesn't exist yet
				mockRepo.On("GetUserByEmail", mock.Anything, "new@example.com").Return(nil, customerror.ErrUserNotFound)
				
				// Role doesn't exist
				mockRepo.On("GetRoleByName", mock.Anything, "user").Return(nil, customerror.ErrNotFound)
			},
			input: &models.CreateUserRequest{
				Username:  "newuser",
				Email:     "new@example.com",
				Password:  "password123",
				FirstName: "New",
				LastName:  "User",
			},
			expectedError: customerror.ErrInternalServer,
		},
		{
			name: "Create User Fails",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// User doesn't exist yet
				mockRepo.On("GetUserByEmail", mock.Anything, "new@example.com").Return(nil, customerror.ErrUserNotFound)
				
				// Role exists
				mockRole := &models.Role{
					ID:   1,
					Name: "user",
				}
				mockRepo.On("GetRoleByName", mock.Anything, "user").Return(mockRole, nil)
				
				// CreateUser fails
				mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(customerror.ErrInternalServer)
			},
			input: &models.CreateUserRequest{
				Username:  "newuser",
				Email:     "new@example.com",
				Password:  "password123",
				FirstName: "New",
				LastName:  "User",
			},
			expectedError: customerror.ErrInternalServer,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock repository
			mockRepo := new(mocks.MockUserRepository)
			
			// Setup mock expectations
			tc.setupMock(mockRepo)
			
			// Create service with mock repository
			service := services.NewAuthService(mockRepo, "test-secret", 24*time.Hour, logger)
			
			// Call the method being tested
			response, err := service.Register(context.Background(), tc.input)
			
			// Assertions
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.True(t, response.Status)
				assert.Contains(t, response.Message, "user created successfully")
				assert.NotNil(t, response.Data)
			}
			
			// Verify all expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}