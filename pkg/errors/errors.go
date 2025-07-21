package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewAppError(code int, message, errorType string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Type:    errorType,
	}
}

// common error types

var (
	ErrInvalidCredentials = NewAppError(http.StatusUnauthorized, "Invalid credentials", "INVALID_CREDENTIALS")
	ErrUnAuthorized       = NewAppError(http.StatusUnauthorized, "Unauthorized access", "UNAUTHORIZED")
	ErrUserNotFound       = NewAppError(http.StatusNotFound, "User not found", "USER_NOT_FOUND")
	ErrUserExists         = NewAppError(http.StatusConflict, "User already exists", "USER_EXISTS")
	ErrInvalidInput       = NewAppError(http.StatusBadRequest, "Invalid input", "INVALID_INPUT")
	ErrInternalServer     = NewAppError(http.StatusInternalServerError, "Internal server error", "INTERNAL")
)


package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Err     error  `json:"-"` // Underlying error, typically not exposed via JSON
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As
func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code int, message, errorType string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Type:    errorType,
	}
}

// NewAppErrorWithCause allows wrapping an underlying error
func NewAppErrorWithCause(code int, message, errorType string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Type:    errorType,
		Err:     err,
	}
}

// common error types (you might still use NewAppError for these)
var (
	ErrInvalidCredentials = NewAppError(http.StatusUnauthorized, "Invalid credentials", "INVALID_CREDENTIALS")
	ErrUnAuthorized       = NewAppError(http.StatusUnauthorized, "Unauthorized access", "UNAUTHORIZED")
	ErrUserNotFound       = NewAppError(http.StatusNotFound, "User not found", "USER_NOT_FOUND")
	ErrUserExists         = NewAppError(http.StatusConflict, "User already exists", "USER_EXISTS")
	ErrInvalidInput       = NewAppError(http.StatusBadRequest, "Invalid input", "INVALID_INPUT")
	ErrInternalServer     = NewAppError(http.StatusInternalServerError, "Internal server error", "INTERNAL")
)