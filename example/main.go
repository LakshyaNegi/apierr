package main

import (
	"apierr"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Creator implements apierr.APIErrorCreator for custom API errors.
type Creator struct{}

func (c *Creator) New() apierr.APIError {
	return NewAPIError()
}

func NewCreator() apierr.APIErrorCreator {
	return &Creator{}
}

// APIError represents a custom API error structure.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Parse maps a CustomError to the APIError structure.
func (e *APIError) FromCustomError(err *apierr.CustomError) error {
	e.Code = err.StatusCode
	e.Message = err.UserMessage
	return nil
}

func NewAPIError() apierr.APIError {
	return &APIError{}
}

func main() {
	// Generate error definitions from YAML (if required)
	err := apierr.Generate("example/errors.yml", "example/errors.gen.go")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create a generic error handler
	creator := NewCreator()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		handler := apierr.NewErrHandler(creator, func() apierr.ResponseWriter {
			return &EchoResponseWriter{ctx: c}
		})
		handler(err)
	}

	// Routes to test various error scenarios
	e.GET("/simple", func(c echo.Context) error {
		return errors.New("simple error")
	})

	e.GET("/complex", func(c echo.Context) error {
		return apierr.NewFromError(
			errors.New("complex error"),
			http.StatusUnauthorized,
			"Unauthorized",
			"AUTH_ERROR",
			"AUTH_401",
			false,
		)
	})

	e.GET("/test", func(c echo.Context) error {
		return apierr.New(
			http.StatusBadRequest,
			"Bad Request",
			"Invalid request format",
			"BAD_REQUEST",
			"BR_400",
			false,
		)
	})

	e.GET("/wrap", func(c echo.Context) error {
		err := apierr.New(
			http.StatusBadRequest,
			"Wrapped Bad Request",
			"Invalid data",
			"BAD_REQUEST",
			"BR_400",
			false,
		)
		return fmt.Errorf("wrap: %w", err)
	})

	e.GET("/double", func(c echo.Context) error {
		err := simulateError()
		return fmt.Errorf("outer error: %w", fmt.Errorf("inner error: %w", err))
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// EchoResponseWriter adapts Echo's context to the apierr.ResponseWriter interface.
type EchoResponseWriter struct {
	ctx echo.Context
}

func (w *EchoResponseWriter) WriteResponse(statusCode int, body interface{}) error {
	return w.ctx.JSON(statusCode, body)
}

// Simulate an error for testing purposes.
func simulateError() error {
	return apierr.New(
		http.StatusNotFound,
		"Resource not found",
		"The requested resource does not exist",
		"NOT_FOUND",
		"NF_404",
		false,
	)
}
