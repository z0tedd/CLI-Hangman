package utils

import (
	"errors"
	"testing"
)

func TestFileOpenError(t *testing.T) {
	var err FileOpenError
	err.Err = errors.New("test")
	if err.Error() != "failed to open file: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}

func TestFileReadError(t *testing.T) {
	var err FileReadError
	err.Err = errors.New("test")
	if err.Error() != "failed to read file: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}

func TestJSONUnmarshalError(t *testing.T) {
	var err JSONUnmarshalError
	err.Err = errors.New("test")
	if err.Error() != "failed to unmarshal JSON: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}
