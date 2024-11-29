package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"uala/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

type GenericError struct {
	Message string `json:"message"`
}

func RespondError(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json;")
	var appErr common.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusInternalServerError, common.AppError{Code: 500, Msg: common.ErrCodeInternalServerError})
	return
}

func RespondErrorBinding(err error, obj interface{}) ValidationErrorResponse {
	switch err := err.(type) {
	case validator.ValidationErrors:
		var errors []ValidationError
		for _, fieldError := range err {
			errors = append(errors, processFieldError(fieldError))
		}
		return ValidationErrorResponse{
			Message: "Validation Errors",
			Code:    http.StatusUnprocessableEntity,
			Errors:  errors,
		}

	case *json.SyntaxError:
		return ValidationErrorResponse{
			Message: fmt.Sprintf("JSON syntax error at byte offset %d", err.Offset),
			Code:    http.StatusInternalServerError,
		}
	default:
		return ValidationErrorResponse{
			Message: fmt.Sprintf("default JSON syntax error at byte offset %s", err.Error()),
			Code:    http.StatusInternalServerError,
		}
	}
}

func processFieldError(fieldError validator.FieldError) ValidationError {
	fullFieldName := fieldError.Namespace()            // Namespace provides the full name including nested struct names
	fieldName := strings.Split(fullFieldName, ".")[1:] // Remove the top-level struct name

	return ValidationError{
		Field:   strings.Join(fieldName, "."), // Rebuild the full field name from the struct hierarchy
		Message: fmt.Sprintf("%s is %s", fieldError.Field(), fieldError.Tag()),
	}
}
