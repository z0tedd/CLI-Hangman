package config_test

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/config"
)

func TestConfig(t *testing.T) {
	Test := config.GetOptions(1)
	if Test.MaxAttempts != 12 {
		t.Errorf("MaxAttempts value isn't right: %d", Test.MaxAttempts)
	}
}
