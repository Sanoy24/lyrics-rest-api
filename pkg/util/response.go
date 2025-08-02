package util

import "github.com/Sanoy24/lyrics-rest-api/internal/models"

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	// Error   any    `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Sucess     bool        `json:"success"`
	Message    string      `json:"message"`
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	HasNext     bool `json:"hasNext"`
	HasPrev     bool `json:"hasPrev"`
	CurrentPage int  `json:"currentPage"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
	Total       int  `json:"total"`
	TotalPages  int  `json:"total_pages"`
}

type AuthResponse struct {
	Token string              `json:"token"`
	User  models.UserResponse `json:"user"`
}

type Response struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorData `json:"error,omitempty"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status bool `json:"status"`
	Error  any  `json:"error"`
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type AppError struct {
	Code    string
	Message string
	Details []FieldError
}
