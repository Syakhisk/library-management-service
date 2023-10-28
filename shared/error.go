package shared

import "errors"

var (
	ErrEmptyBody = errors.New("body can not be empty")
	ErrCreateBook = errors.New("failed to create book")
)
