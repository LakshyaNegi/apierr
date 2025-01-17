package apierr

import (
	"fmt"
)

const (
	// Error Names
	ErrorBadRequest          = "BadRequest"          // The request is invalid or malformed.
	ErrorUnauthorized        = "Unauthorized"        // The user is not authorized to perform this action.
	ErrorForbidden           = "Forbidden"           // The user does not have permission to access this resource.
	ErrorNotFound            = "NotFound"            // The requested resource could not be found.
	ErrorInternalServerError = "InternalServerError" // An unexpected server error occurred.
	ErrorParseError          = "ParseError"          // Failed to parse API Error.

	// Error Types
	ErrTypeBAD_REQUEST           = "BAD_REQUEST"
	ErrTypeUNAUTHORIZED          = "UNAUTHORIZED"
	ErrTypeFORBIDDEN             = "FORBIDDEN"
	ErrTypeNOT_FOUND             = "NOT_FOUND"
	ErrTypeINTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	ErrTypePARSE_ERROR           = "PARSE_ERROR"

	// Error Codes
	ErrCodeBAD_REQUEST           = "BAD_REQUEST"
	ErrCodeUNAUTHORIZED          = "UNAUTHORIZED"
	ErrCodeFORBIDDEN             = "FORBIDDEN"
	ErrCodeNOT_FOUND             = "NOT_FOUND"
	ErrCodeINTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	ErrCodePARSE_ERROR           = "PARSE_ERROR"
)

// BadRequestError represents The request is invalid or malformed..
type BadRequestError struct {
	CustomError
	details string
}

// NewBadRequestError creates a new BadRequestError.
func NewBadRequestError(details string,
) *CustomError {
	return New(
		400,
		fmt.Sprintf(
			"Invalid request: %s",
			details,
		),
		"The request is invalid.",
		"BAD_REQUEST",
		"BAD_REQUEST",
		false,
	)
}

// GetDetails returns the value of details for BadRequestError.
func (e *BadRequestError) GetDetails() string {
	return e.details
}

// UnauthorizedError represents The user is not authorized to perform this action..
type UnauthorizedError struct {
	CustomError
}

// NewUnauthorizedError creates a new UnauthorizedError.
func NewUnauthorizedError() *CustomError {
	return New(
		401,
		"Unauthorized access",
		"You are not authorized to access this resource.",
		"UNAUTHORIZED",
		"UNAUTHORIZED",
		false,
	)
}

// ForbiddenError represents The user does not have permission to access this resource..
type ForbiddenError struct {
	CustomError
	resource string
}

// NewForbiddenError creates a new ForbiddenError.
func NewForbiddenError(resource string,
) *CustomError {
	return New(
		403,
		fmt.Sprintf(
			"Forbidden: You do not have permission to access %s.",
			resource,
		),
		"You do not have permission to access this resource.",
		"FORBIDDEN",
		"FORBIDDEN",
		false,
	)
}

// GetResource returns the value of resource for ForbiddenError.
func (e *ForbiddenError) GetResource() string {
	return e.resource
}

// NotFoundError represents The requested resource could not be found..
type NotFoundError struct {
	CustomError
	resource string
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(resource string,
) *CustomError {
	return New(
		404,
		fmt.Sprintf(
			"%s not found",
			resource,
		),
		"The requested resource could not be found.",
		"NOT_FOUND",
		"NOT_FOUND",
		false,
	)
}

// GetResource returns the value of resource for NotFoundError.
func (e *NotFoundError) GetResource() string {
	return e.resource
}

// InternalServerErrorError represents An unexpected server error occurred..
type InternalServerErrorError struct {
	CustomError
}

// NewInternalServerErrorError creates a new InternalServerErrorError.
func NewInternalServerErrorError() *CustomError {
	return New(
		500,
		"Internal server error",
		"Something went wrong. Please try again later.",
		"INTERNAL_SERVER_ERROR",
		"INTERNAL_SERVER_ERROR",
		false,
	)
}

// ParseErrorError represents Failed to parse API Error..
type ParseErrorError struct {
	CustomError
	api_error any
	error     any
}

// NewParseErrorError creates a new ParseErrorError.
func NewParseErrorError(api_error any, error any,
) *CustomError {
	return New(
		500,
		fmt.Sprintf(
			"failed to api_error :%v, parse error :%v",
			api_error, error,
		),
		"Something went wrong. Please try again later.",
		"PARSE_ERROR",
		"PARSE_ERROR",
		false,
	)
}

// GetApi_error returns the value of api_error for ParseErrorError.
func (e *ParseErrorError) GetApiError() any {
	return e.api_error
}

// GetError returns the value of error for ParseErrorError.
func (e *ParseErrorError) GetError() any {
	return e.error
}
