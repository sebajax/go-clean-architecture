package apperror

import (
	"fmt"
	"net/http"
)

// interface to inject into services
type IAppError interface {
	Error() string
	NotFound(message string)
	BadRequest(message string)
	Unauthorized(message string)
	Forbidden(message string)
	InternalServerError()
}

type AppError struct {
	Code    int
	Message string
}

func NewAppError(code int, message string) IAppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (r *AppError) Error() string {
	return fmt.Sprintf(r.Message)
}

func (r *AppError) NotFound(message string) {
	r.Code = http.StatusNotFound
	r.Message = message
}

func (r *AppError) BadRequest(message string) {
	r.Code = http.StatusBadRequest
	r.Message = message
}

func (r *AppError) Unauthorized(message string) {
	r.Code = http.StatusUnauthorized
	r.Message = message
}

func (r *AppError) Forbidden(message string) {
	r.Code = http.StatusForbidden
	r.Message = message
}

func (r *AppError) InternalServerError() {
	r.Code = http.StatusInternalServerError
	r.Message = "INTERNAL_SERVER_ERROR"
}
