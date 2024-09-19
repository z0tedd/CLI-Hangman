package utils

import (
	"fmt"
)

// Custom error types
type FileOpenError struct {
	Err error
}

func (e *FileOpenError) Error() string {
	return fmt.Sprintf("failed to open file: %v", e.Err)
}

type FileReadError struct {
	Err error
}

func (e *FileReadError) Error() string {
	return fmt.Sprintf("failed to read file: %v", e.Err)
}

type JSONUnmarshalError struct {
	Err error
}

func (e *JSONUnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal JSON: %v", e.Err)
}
