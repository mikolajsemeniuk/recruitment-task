package index

import "errors"

var (
	ErrIndexNotFound = errors.New("index not found")
	ErrValueNegative = errors.New("value cannot be negative")
	// Add more errors below...
)
