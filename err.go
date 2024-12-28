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
	err         error
	statusCode  int
	message     string
	userMessage string
	errType     string
	errCode     string
}

// StatusCode returns the HTTP status code.
func (e *CustomError) StatusCode() int {
	return e.statusCode
}

// Message returns the internal error message.
func (e *CustomError) Message() string {
	return e.message
}

// UserMessage returns the user-facing error message.
func (e *CustomError) UserMessage() string {
	return e.userMessage
}

// ErrType returns the type of the error.
func (e *CustomError) ErrType() string {
	return e.errType
}

// ErrCode returns the error code.
func (e *CustomError) ErrCode() string {
	return e.errCode
}

// Error implements the error interface.
func (e *CustomError) Error() string {
	return e.message
}

// Unwrap provides access to the wrapped error.
func (e *CustomError) Unwrap() error {
	return e.err
}

// New creates a new CustomError.
func New(statusCode int, message, userMessage, errType, errCode string) *CustomError {
	return &CustomError{
		err:         fmt.Errorf("custom error: %s", message),
		statusCode:  statusCode,
		message:     message,
		userMessage: userMessage,
		errType:     errType,
		errCode:     errCode,
	}
}

// NewFromError creates a new CustomError from an existing error.
func NewFromError(err error, statusCode int, userMessage, errType, errCode string) *CustomError {
	return &CustomError{
		err:         fmt.Errorf("wrapped error: %w", err),
		statusCode:  statusCode,
		message:     err.Error(),
		userMessage: userMessage,
		errType:     errType,
		errCode:     errCode,
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

			_ = writer.WriteResponse(customErr.StatusCode(), apiErr)
			return
		}

		_ = writer.WriteResponse(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}
}
