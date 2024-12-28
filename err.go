package apierr

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// APIError defines the interface for custom API errors.
type APIError interface {
	FromCustomError(customErr *CustomError) error
}

// APIErrorCreator defines the interface for creating custom API errors.
type APIErrorCreator interface {
	New() APIError
}

// ResponseWriter abstracts response writing for different frameworks.
type ResponseWriter interface {
	WriteResponse(statusCode int, body interface{}) error
}

// CustomError represents a standardized error structure.
type CustomError struct {
	BaseErr     error
	StatusCode  int
	Message     string
	UserMessage string
	ErrType     string
	ErrCode     string
}

// Error implements the error interface.
func (e *CustomError) Error() string {
	return e.Message
}

// Unwrap provides access to the wrapped error.
func (e *CustomError) Unwrap() error {
	return e.BaseErr
}

// New creates a new CustomError.
func New(statusCode int, message, userMessage, errType, errCode string) *CustomError {
	return &CustomError{
		BaseErr:     fmt.Errorf("error: %s", message),
		StatusCode:  statusCode,
		Message:     message,
		UserMessage: userMessage,
		ErrType:     errType,
		ErrCode:     errCode,
	}
}

// NewFromError creates a new CustomError from an existing error.
func NewFromError(err error, statusCode int, userMessage, errType, errCode string) *CustomError {
	return &CustomError{
		BaseErr:     err,
		StatusCode:  statusCode,
		Message:     err.Error(),
		UserMessage: userMessage,
		ErrType:     errType,
		ErrCode:     errCode,
	}
}

// NewErrHandler creates a generic error handler.
func NewErrHandler(creator APIErrorCreator, writerFactory func() ResponseWriter) func(error) {
	return func(werr error) {
		var customErr *CustomError
		writer := writerFactory()

		if errors.As(werr, &customErr) {
			apiErr := creator.New()
			if err := apiErr.FromCustomError(customErr); err != nil {
				slog.Warn(
					"failed to process error",
					"error", err,
					"argument error", werr,
				)

				_ = writer.WriteResponse(http.StatusInternalServerError, map[string]string{
					"error": "Failed to process error.",
				})
				return
			}

			_ = writer.WriteResponse(customErr.StatusCode, apiErr)
			return
		}

		_ = writer.WriteResponse(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}
}
