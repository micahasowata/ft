package ft

import "errors"

var (
	ErrNotFound      = errors.New("extension not found")
	ErrInvalidFormat = errors.New("extension does not begin with a dot")
)
