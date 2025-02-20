package model

import "errors"

var ErrNameRequired error = errors.New("name is required")
var ErrEmailRequired error = errors.New("email is required")
var ErrEmailInvalid error = errors.New("email is invalid")
var ErrInvalidUUID = errors.New("invalid user id")
var ErrRequiredUserID = errors.New("user id is required")
