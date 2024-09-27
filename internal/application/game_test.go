package application_test

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	var got *domain.Game

	variableAttempts := []int{12, 9, 7, 7}

	for i := range len(variableAttempts) {
		got = application.NewGame("apple", i+1)
		assert.Equal(t, got.AttemptsLeft, variableAttempts[i], "got - %q", got.AttemptsLeft)
	}
}
