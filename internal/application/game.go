package application

import (
	"strings"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/config"
	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
)

func NewGame(word string, difficulty int) *domain.Game {
	options := config.GetOptions(difficulty)

	maxAttempts := options.MaxAttempts

	return &domain.Game{
		UsedLetters:  make(map[rune]bool),
		WordToGuess:  strings.ToUpper(word),
		Guessed:      []rune(strings.Repeat("~", len(word))),
		AttemptsLeft: maxAttempts,
		MaxAttempts:  maxAttempts,
		Difficulty:   difficulty,
	}
}
