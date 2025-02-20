package errhandler

import "fmt"

type AppError struct {
	Code    string
	Message string
}

type ApplicationError struct {
	AppError
}
type NotFoundError struct {
	AppError
}
type ValidationError struct {
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

func NewNotFoundError(message string) error {
	return &NotFoundError{
		AppError: AppError{
			Code:    "NOT_FOUND",
			Message: message,
		},
	}
}
