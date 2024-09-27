package domain_test

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	usedLetters = map[rune]bool{'A': true}
	g           = domain.Game{
		UsedLetters:  usedLetters,
		WordToGuess:  "PAAA",
		Guessed:      []rune{'~', '~', 'A', '~'},
		AttemptsLeft: 7,
		MaxAttempts:  7,
		Difficulty:   1,
	}
)

func TestUpdateGuessedWord(t *testing.T) {
	g.UpdateGuessedWord('P')
	assert.Equal(t, "P~A~", string(g.Guessed), "Guessed word isn't matched: expected - %s, got - %b", "P~A~", string(g.Guessed))
}

func TestIsWordGuessed(t *testing.T) {
	assert.Equal(t, false, g.IsWordGuessed(), "expected - true, got - %b", g.IsWordGuessed())
}

func TestEndGame(t *testing.T) {
	assert.Equal(t, "Вы проиграли! Загаданное слово было: PAAA\n", g.EndGame(), "EndGame doesn't work right\n got: %s", g.EndGame())
}

func TestIsLetterInWord(t *testing.T) {
	assert.Equal(t, g.IsLetterInWord('A'), true, "IsletterInWord must be true")
}

func TestIsNotWordGuessed(t *testing.T) {
	g.UpdateGuessedWord('A')
	assert.Equal(t, true, g.IsWordGuessed(), "expected - true, got - %b", g.IsWordGuessed())
}
