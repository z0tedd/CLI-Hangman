package domain

import (
	"testing"
)

var (
	usedLetters = map[rune]bool{'A': true}
	g           = Game{
		UsedLetters:  usedLetters,
		WordToGuess:  "PAAA",
		Guessed:      []rune{'~', '~', 'A', '~'},
		AttemptsLeft: 7,
		MaxAttempts:  7,
		Difficulty:   1,
	}
)

func TestUpdateGuessedWord(t *testing.T) {
	g.updateGuessedWord('P')
	if string(g.Guessed) != "P~A~" {
		t.Error("Guessed word isn't matched")
	}
}

func TestIsWordGuessed(t *testing.T) {
	if g.isWordGuessed() == true {
		t.Error("function isWordGuessed doesn't work right")
	}
}

func TestEndGame(t *testing.T) {
	if g.endGame() != "Вы проиграли! Загаданное слово было: PAAA\n" {
		t.Errorf("EndGame doesn't work right\n got: %s", g.endGame())
	}
}

func TestIsLetterInWord(t *testing.T) {
	t.Log(g.AttemptsLeft)
	if g.IsLetterInWord('A') != true {
		t.Errorf("s.Display() = %q, want %q", 1, 0)
	}
}

func TestIsNotWordGuessed(t *testing.T) {
	g.updateGuessedWord('A')
	if g.isWordGuessed() == false {
		t.Error("function isWordGuessed doesn't work right")
	}
}
