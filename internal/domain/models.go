package domain

import (
	"fmt"
	"os"
	"strings"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/infrastructure"
)

type Game struct {
	UsedLetters  map[rune]bool
	WordToGuess  string
	Guessed      []rune
	AttemptsLeft int
	MaxAttempts  int
	Difficulty   int
}

func (g *Game) Start() {
	for g.AttemptsLeft > 0 && !g.isWordGuessed() {
		infrastructure.DisplayGameState(os.Stdout, g.Guessed, g.AttemptsLeft, g.MaxAttempts)

		var input string
		fmt.Print("Введите букву: ")
		fmt.Scanln(&input)
		input = strings.ToUpper(input)

		if len(input) != 1 {
			fmt.Println("Введите одну букву!")
			continue
		}

		letter := rune(input[0])
		if g.UsedLetters[letter] {
			fmt.Println("Вы уже вводили эту букву.")
			continue
		}

		g.UsedLetters[letter] = true
		if g.IsLetterInWord(letter) {
			g.updateGuessedWord(letter)
		} else {
			g.AttemptsLeft--
		}
	}
	infrastructure.DisplayGameState(os.Stdout, g.Guessed, g.AttemptsLeft, g.MaxAttempts)
	fmt.Println(g.endGame())
}

func (g *Game) IsLetterInWord(letter rune) bool {
	return strings.ContainsRune(g.WordToGuess, letter)
}

func (g *Game) updateGuessedWord(letter rune) {
	for i, l := range g.WordToGuess {
		if l == letter {
			g.Guessed[i] = letter
		}
	}
}

func (g *Game) isWordGuessed() bool {
	for _, letter := range g.Guessed {
		if letter == '~' {
			return false
		}
	}
	return true
}

func (g *Game) endGame() string {
	if g.isWordGuessed() {
		return fmt.Sprintln("Вы угадали слово:", g.WordToGuess)
	}
	return fmt.Sprintln("Вы проиграли! Загаданное слово было:", g.WordToGuess)
}
