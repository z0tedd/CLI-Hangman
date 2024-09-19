package domain_test

import (
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
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

	if string(g.Guessed) != "P~A~" {
		t.Error("Guessed word isn't matched")
	}
}

func TestIsWordGuessed(t *testing.T) {
	if g.IsWordGuessed() == true {
		t.Error("function IsWordGuessed doesn't work right")
	}
}

func TestEndGame(t *testing.T) {
	if g.EndGame() != "Вы проиграли! Загаданное слово было: PAAA\n" {
		t.Errorf("EndGame doesn't work right\n got: %s", g.EndGame())
	}
}

func TestIsLetterInWord(t *testing.T) {
	if g.IsLetterInWord('A') != true {
		t.Errorf("s.Display() = %q, want %q", 1, 0)
	}
}

func TestIsNotWordGuessed(t *testing.T) {
	g.UpdateGuessedWord('A')

	if g.IsWordGuessed() == false {
		t.Error("function IsWordGuessed doesn't work right")
	}
}
