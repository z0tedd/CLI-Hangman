package infrastructure_test

import (
	"bytes"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestDisplayGameState(t *testing.T) {
	var b bytes.Buffer

	guessed := []rune("~~~~~~")
	attemtsleft := 12
	maxattempts := 12

	infrastructure.DisplayGameState(&b, guessed, attemtsleft, maxattempts)
	got := b.String()
	want := `
Слово:  ~~~~~~

     _______
     |    \|
           |
           |
           |
          /|
    =========
Осталось попыток: 12

`
	assert.Equal(t, got, want, "got %q, want %q", got, want)

	b.Reset()

	infrastructure.DisplayGameState(&b, guessed, 0, maxattempts)
	got = b.String()
	want = `
Слово:  ~~~~~~

     _______
     |    \|
     O     |
    /|\    |
    / \    |
          /|
    =========
Осталось попыток: 0

`

	assert.Equal(t, got, want, "got %q, want %q", got, want)
}
