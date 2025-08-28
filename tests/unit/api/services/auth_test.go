package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces/mocks"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/Sanoy24/lyrics-rest-api/tests/unit/api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestAuthService_Login(t *testing.T) {
	// Setup logger
	logger, _ := zap.NewDevelopment()

	// Test cases
	tests := []struct {
		name          string
		setupMock     func(mockRepo *mocks.MockUserRepository)
		input         *models.UserLoginRequest
		expectedError error
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// Hash a known password for testing
				hashedPassword, _ := util.HashPassword("password123")
				
				// Create a mock user with the hashed password
				mockUser := testutils.CreateTestUser()
				mockUser.Password = hashedPassword
				
				mockRepo.On("GetByUsernameOrPassword", mock.Anything, "testuser").Return(mockUser, nil)
			},
			input: &models.UserLoginRequest{
				Identifier: "testuser",
				Password:   "password123",
			},
			expectedError: nil,
		},
		{
			name: "User Not Found",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.On("GetByUsernameOrPassword", mock.Anything, "nonexistent").Return(nil, customerror.ErrUserNotFound)
			},
			input: &models.UserLoginRequest{
				Identifier: "nonexistent",
				Password:   "password123",
			},
			expectedError: customerror.ErrInvalidCredentials,
		},
		{
			name: "Wrong Password",
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				// Hash a different password than what will be provided
				hashedPassword, _ := util.HashPassword("correctpassword")
				
				mockUser := testutils.CreateTestUser()
				mockUser.Password = hashedPassword
				
				mockRepo.On("GetByUsernameOrPassword", mock.Anything, "testuser").Return(mockUser, nil)
			},
			input: &models.UserLoginRequest{
				Identifier: "testuser",
				Password:   "wrongpassword",
			},
			expectedError: customerror.ErrInvalidCredentials,
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
			response, err := service.Login(context.Background(), tc.input)
			
			// Assertions
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.Token)
				assert.Equal(t, uint(1), response.User.ID)
				assert.Equal(t, "testuser", response.User.Username)
			}
			
			// Verify all expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}