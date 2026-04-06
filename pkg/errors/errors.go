package errs

import "errors"

// Error type
type AppError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

// Authentication errors
var (
	ErrMissingAuthHeader = errors.New("authorization header is required")
	ErrWrongAuthHeader   = errors.New("authorization header must start with 'API-KEY '")
	ErrInvalidAPIKey     = errors.New("invalid API key")
)
