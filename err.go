package apierr

import (
	"errors"
	"fmt"
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
	BaseErr     error                  // Base error
	StatusCode  int                    // HTTP status code
	Message     string                 // Error message
	UserMessage string                 // User-friendly error message
	ErrType     string                 // Error type
	ErrCode     string                 // Unique error code
	Retryable   bool                   // Indicates if the error is retryable
	Metadata    map[string]interface{} // Additional contextual information
}

// Error implements the error interface.
func (e *CustomError) Error() string {
	return e.Message
}

// Unwrap provides access to the wrapped error.
func (e *CustomError) Unwrap() error {
	return e.BaseErr
}

// SetMetadata sets a metadata key-value pair.
func (e *CustomError) SetMetadata(key string, value interface{}) {
	if e.Metadata == nil {
		e.Metadata = make(map[string]interface{})
	}
	e.Metadata[key] = value
}

// GetMetadata returns the value of a metadata key.
func (e *CustomError) GetMetadata(key string) (interface{}, bool) {
	if e.Metadata == nil {
		return nil, false
	}
	value, ok := e.Metadata[key]
	return value, ok
}

// New creates a new CustomError.
func New(statusCode int, message, userMessage, errType, errCode string, retryable bool) *CustomError {
	return &CustomError{
		BaseErr:     fmt.Errorf("error: %s", message),
		StatusCode:  statusCode,
		Message:     message,
		UserMessage: userMessage,
		ErrType:     errType,
		ErrCode:     errCode,
		Retryable:   retryable,
		Metadata:    make(map[string]interface{}), // Initialize an empty metadata map
	}
}

// NewFromError creates a new CustomError from an existing error.
func NewFromError(err error, statusCode int, userMessage, errType, errCode string, retryable bool) *CustomError {
	return &CustomError{
		BaseErr:     err,
		StatusCode:  statusCode,
		Message:     err.Error(),
		UserMessage: userMessage,
		ErrType:     errType,
		ErrCode:     errCode,
		Retryable:   retryable,
		Metadata:    make(map[string]interface{}), // Initialize an empty metadata map
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
				_ = writer.WriteResponse(http.StatusInternalServerError, NewParseErrorError(werr, err))
				return
			}

			_ = writer.WriteResponse(customErr.StatusCode, apiErr)
			return
		}

		_ = writer.WriteResponse(http.StatusInternalServerError, NewInternalServerErrorError())
	}
}
