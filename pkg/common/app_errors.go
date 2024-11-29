package common

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicateKey        = "duplicate_key"
	ErrCodeInternalServerError = "internal server error"
	ErrCodeInvalidParams       = "invalid_params"
	ErrCodeNotFound            = 404
	ErrCodeTimeout             = "timeout"
	ErrCodeInternalServer      = 500
)

var (
	ErrDuplicateKey = errors.New("duplicate key error")
	ErrNotFound     = errors.New("record not found error")
	ErrTimeout      = errors.New("timeout error")
	ErridNotValid   = errors.New("id is not valid")
	ErrCreate       = errors.New("error creating record")
	ErrUpdate       = errors.New("error updating record")
	ErrDelete       = errors.New("error deleting record")
	ErrRetrieve     = errors.New("error retrieving record")
)

// AppError is a custom error type that implements the error interface
type AppError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// NewAppError creates a new AppError with the given code and message.
func NewAppError(code int, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

// Error returns a string representation of the error. It is part of the error interface.
func (e AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
