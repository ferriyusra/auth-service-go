package myerror

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("not found")
	ErrInvalidFilter = errors.New("invalid filter")
	ErrValidation    = errors.New("validation error")
)
