package customerror

import "fmt"

type AppError struct {
	Code    string
	Message string
}

type ApplicationError struct {
	AppError
}

type TimeoutError struct {
	AppError
}
type NotFoundError struct {
	AppError
}
type ValidationError struct {
	AppError
}

type UnauthorizedError struct {
	AppError
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewValidationError(message string) error {
	return &ValidationError{
		AppError: AppError{
			Code:    "VALIDATION_ERROR",
			Message: message,
		},
	}
}

func NewApplicationError(message string) error {
	return &ApplicationError{
		AppError: AppError{
			Code:    "APPLICATION_ERROR",
			Message: message,
		},
	}
}

func NewTimeoutError(message string) error {
	return &TimeoutError{
		AppError: AppError{
			Code:    "TIMEOUT_ERROR",
			Message: message,
		},
	}
}

func NewNotFoundError(message string) error {
	return &NotFoundError{
		AppError: AppError{
			Code:    "NOT_FOUND",
			Message: message,
		},
	}
}

func NewUnauthorizedError(message string) error {
	return &UnauthorizedError{
		AppError: AppError{
			Code:    "AUTHORIZATION_ERROR",
			Message: message,
		},
	}
}
