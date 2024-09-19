package application

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
)

func TestNewGame(t *testing.T) {
	var got *domain.Game
	variableAttempts := []int{12, 9, 7, 7}
	for i := range len(variableAttempts) {
		got = NewGame("apple", i+1)

		if got.AttemptsLeft != variableAttempts[i] {
			t.Fatalf("got - %q", got.AttemptsLeft)
		}
	}
}
