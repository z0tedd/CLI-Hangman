package application

import (
	"strings"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
)

func NewGame(word string, difficulty int) *domain.Game {
	var maxAttempts int
	var guessedword []rune
	switch difficulty {
	case 1:
		maxAttempts = 12
	case 2:
		maxAttempts = 9
	case 3:
		maxAttempts = 7
	default:
		maxAttempts = 7
	}

	guessedword = make([]rune, len(word))
	for i := range word {
		guessedword[i] = '~'
	}
	// fmt.Println(guessedword)
	return &domain.Game{
		UsedLetters:  make(map[rune]bool),
		WordToGuess:  strings.ToUpper(word),
		Guessed:      guessedword,
		AttemptsLeft: maxAttempts,
		MaxAttempts:  maxAttempts,
		Difficulty:   difficulty,
	}
}
