package config_test

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	testcases := []struct {
		text  string
		input int
		want  int
	}{
		// Table
		{"difficulty is 1", 1, 12},
		{"difficulty is 2", 2, 9},
		{"difficulty is 3", 3, 7},
		{"difficulty is not in options", 10, 7},
	}
	for _, tt := range testcases {
		t.Run(tt.text, func(t *testing.T) {
			result := config.GetOptions(tt.input).MaxAttempts
			assert.Equal(t, result, tt.want, "got %d, want %d", result, tt.want)
		})
	}
}
