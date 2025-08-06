package customerror

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserExists         = NewAppError("USER_EXISTS", "User with this email or username already exists", http.StatusConflict)
	ErrUserNotFound       = NewAppError("USER_NOT_FOUND", "User not found", http.StatusNotFound)
	ErrInvalidInput       = NewAppError("INVALID_INPUT", "Invalid input data", http.StatusBadRequest)
	ErrInternalServer     = NewAppError("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
	ErrUnauthorized       = NewAppError("UNAUTHORIZED", "Unauthorized access", http.StatusUnauthorized)
	ErrForbidden          = NewAppError("FORBIDDEN", "Access forbidden", http.StatusForbidden)
	ErrInvalidCredentials = NewAppError("INVALID_CREDENTIALS", "Invalid credentials", http.StatusUnauthorized)
	ErrNotFound           = errors.New("resource not found")
)

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewValidationError(field, message string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    fmt.Sprintf("Field '%s': %s", field, message),
		StatusCode: http.StatusBadRequest,
	}
}

func IsDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		return pgError.Code == "23505" // Unique violation
	}
	errStr := strings.ToLower(err.Error())
	duplicateKeywords := []string{
		"duplicate key value violates unique constraint",
		"unique constraint",
		"duplicate entry",
		"already exists",
	}
	for _, keyword := range duplicateKeywords {
		if strings.Contains(errStr, keyword) {
			return true
		}
	}
	return false

}

func GetDuplicateKeyErrorField(err error) string {
	if err == nil {
		return ""
	}
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) && pgError.Code == "23505" {
		errorDetail := pgError.Detail
		if strings.Contains(errorDetail, "(email)") {
			return "email"
		}
		if strings.Contains(errorDetail, "(username)") {
			return "username"
		}

		constraintName := pgError.ConstraintName
		if strings.Contains(constraintName, "email") {
			return "email"
		}
		if strings.Contains(constraintName, "username") {
			return "username"
		}
	}
	return ""
}
