package utils_test

import (
	"errors"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFileOpenError(t *testing.T) {
	var err utils.FileOpenError

	err.Err = errors.New("test")
	assert.Equal(t, err.Error(), "failed to open file: test", "error is failed - %q", err.Err)
}

func TestFileReadError(t *testing.T) {
	var err utils.FileReadError

	err.Err = errors.New("test")
	assert.Equal(t, err.Error(), "failed to read file: test", "error is failed - %q", err.Err)
}

func TestJSONUnmarshalError(t *testing.T) {
	var err utils.JSONUnmarshalError

	err.Err = errors.New("test")
	assert.Equal(t, err.Error(), "failed to unmarshal JSON: test", "error is failed - %q", err.Err)
}
