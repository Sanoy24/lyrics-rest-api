package services_test

import (
	"context"
	"testing"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces/mocks"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
	"github.com/Sanoy24/lyrics-rest-api/tests/unit/api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestGetCurrentUser(t *testing.T) {
	tests := []struct {
		name      string
		userID    int
		setupMock func(*mocks.MockUserRepository)
		wantErr   bool
		errType   error
	}{
		{
			name:   "Success - User found",
			userID: 1,
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				testUser := testutils.CreateTestUser()
				mockRepo.On("GetUserByID", mock.Anything, 1).Return(testUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "Error - User not found",
			userID: 999,
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.On("GetUserByID", mock.Anything, 999).Return(nil, customerror.ErrUserNotFound)
			},
			wantErr: true,
			errType: customerror.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockRepo := new(mocks.MockUserRepository)
			tt.setupMock(mockRepo)

			// Create logger
			logger, _ := zap.NewDevelopment()

			// Create service
			userService := services.NewUserService(mockRepo, logger)

			// Call the method
			user, err := userService.GetCurrentUser(context.Background(), tt.userID)

			// Assert results
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, uint(tt.userID), user.ID)
			}

			// Verify all expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name      string
		userID    int
		request   *models.UpdateUserRequest
		setupMock func(*mocks.MockUserRepository)
		wantErr   bool
		errType   error
	}{
		{
			name:   "Success - User updated",
			userID: 1,
			request: &models.UpdateUserRequest{
				FirstName: "Updated",
				LastName:  "User",
				Bio:       "Updated bio",
			},
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.On("UpdateUser", mock.Anything, 1, mock.AnythingOfType("*models.UpdateUserRequest")).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Error - User not found",
			userID: 999,
			request: &models.UpdateUserRequest{
				FirstName: "Updated",
				LastName:  "User",
				Bio:       "Updated bio",
			},
			setupMock: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.On("UpdateUser", mock.Anything, 999, mock.AnythingOfType("*models.UpdateUserRequest")).Return(customerror.ErrUserNotFound)
			},
			wantErr: true,
			errType: customerror.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockRepo := new(mocks.MockUserRepository)
			tt.setupMock(mockRepo)

			// Create logger
			logger, _ := zap.NewDevelopment()

			// Create service
			userService := services.NewUserService(mockRepo, logger)

			// Call the method
			err := userService.UpdateUser(context.Background(), tt.userID, tt.request)

			// Assert results
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}