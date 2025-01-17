// Code generated by error-gen. DO NOT EDIT.

package errExample

import (
	"github.com/LakshyaNegi/apierr"
	"fmt"
)

const (
	// Error Names
	ErrorBadRequest = "BadRequest" // The request is invalid or malformed.
	ErrorUnauthorized = "Unauthorized" // The user is not authorized to perform this action.
	ErrorForbidden = "Forbidden" // The user does not have permission to access this resource.
	ErrorNotFound = "NotFound" // The requested resource could not be found.
	ErrorInternalServerError = "InternalServerError" // An unexpected server error occurred.
	ErrorParseError = "ParseError" // Failed to parse API Error.

	// Error Types
	ErrTypeBadRequest = "BAD_REQUEST"
	ErrTypeUnauthorized = "UNAUTHORIZED"
	ErrTypeForbidden = "FORBIDDEN"
	ErrTypeNotFound = "NOT_FOUND"
	ErrTypeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrTypeParseError = "PARSE_ERROR"

	// Error Codes
	ErrCodeBadRequest = "BAD_REQUEST"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden = "FORBIDDEN"
	ErrCodeNotFound = "NOT_FOUND"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeParseError = "PARSE_ERROR"
)


// BadRequestError represents The request is invalid or malformed..
type BadRequestError struct {
	apierr.CustomError
	details string
}

// NewBadRequestError creates a new BadRequestError.
func NewBadRequestError(details string,
) *apierr.CustomError {
	return apierr.New(
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
	apierr.CustomError
}

// NewUnauthorizedError creates a new UnauthorizedError.
func NewUnauthorizedError(
) *apierr.CustomError {
	return apierr.New(
		401,
		fmt.Sprintf(
			"Unauthorized access",
			
		),
		"You are not authorized to access this resource.",
		"UNAUTHORIZED",
		"UNAUTHORIZED",
		false,
	)
}

// ForbiddenError represents The user does not have permission to access this resource..
type ForbiddenError struct {
	apierr.CustomError
	resource string
}

// NewForbiddenError creates a new ForbiddenError.
func NewForbiddenError(resource string,
) *apierr.CustomError {
	return apierr.New(
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
	apierr.CustomError
	resource string
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(resource string,
) *apierr.CustomError {
	return apierr.New(
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
	apierr.CustomError
}

// NewInternalServerErrorError creates a new InternalServerErrorError.
func NewInternalServerErrorError(
) *apierr.CustomError {
	return apierr.New(
		500,
		fmt.Sprintf(
			"Internal server error",
			
		),
		"Something went wrong. Please try again later.",
		"INTERNAL_SERVER_ERROR",
		"INTERNAL_SERVER_ERROR",
		false,
	)
}

// ParseErrorError represents Failed to parse API Error..
type ParseErrorError struct {
	apierr.CustomError
	apiError any
	err error
}

// NewParseErrorError creates a new ParseErrorError.
func NewParseErrorError(apiError any,err error,
) *apierr.CustomError {
	return apierr.New(
		500,
		fmt.Sprintf(
			"failed to parse API error :%v, parse error :%v",
			apiError,err,
		),
		"Something went wrong. Please try again later.",
		"PARSE_ERROR",
		"PARSE_ERROR",
		false,
	)
}

// GetApierror returns the value of apiError for ParseErrorError.
func (e *ParseErrorError) GetApierror() any {
	return e.apiError
}

// GetErr returns the value of err for ParseErrorError.
func (e *ParseErrorError) GetErr() error {
	return e.err
}

