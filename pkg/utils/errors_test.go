package utils_test

import (
	"errors"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
)

func TestFileOpenError(t *testing.T) {
	var err utils.FileOpenError

	err.Err = errors.New("test")
	if err.Error() != "failed to open file: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}

func TestFileReadError(t *testing.T) {
	var err utils.FileReadError

	err.Err = errors.New("test")
	if err.Error() != "failed to read file: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}

func TestJSONUnmarshalError(t *testing.T) {
	var err utils.JSONUnmarshalError

	err.Err = errors.New("test")
	if err.Error() != "failed to unmarshal JSON: test" {
		t.Fatalf("error is failed - %q", err.Err)
	}
}
