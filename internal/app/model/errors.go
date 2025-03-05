package model

import "errors"

var ErrInvalidRequest = errors.New("request malformed or invalid")
var ErrRequiredName error = errors.New("name is required")
var ErrInvalidName error = errors.New("name max length are 120 characters")
var ErrRequiredEmail error = errors.New("email is required")
var ErrInvalidEmail error = errors.New("email is invalid")
var ErrInvalidUUID = errors.New("invalid user id")
var ErrRequiredUserID = errors.New("user id is required")
var ErrRequiredPassword = errors.New("password is required")
var ErrInvalidPasswordLength = errors.New("password should be of length between 8 and 64 characters")
var ErrInvalidPassword = errors.New("password require a combination of uppercase, lowercase, numbers, and special characters")
var ErrRequiredCredentials = errors.New("credentials are required")
var ErrInvalidCredentials = errors.New("credentials are invalid")
var ErrInvalidGrantType = errors.New("invalid grant type: supported types are 'client_credentials'")
