package domain

import "github.com/pkg/errors"

// ErrInvalidInput Invalid Input Error
var ErrInvalidInput = errors.New("Invalid Input")

// ErrNotFound Not Found Error
var ErrNotFound = errors.New("Not Found")

// ErrAuth Authorization Error
var ErrAuth = errors.New("Not Authorized to access resource")
