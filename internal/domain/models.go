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
	for g.AttemptsLeft > 0 && !g.IsWordGuessed() {
		infrastructure.DisplayGameState(os.Stdout, g.Guessed, g.AttemptsLeft, g.MaxAttempts)

		var input string

		fmt.Print("Введите букву: ")

		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}

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
			g.UpdateGuessedWord(letter)
		} else {
			g.AttemptsLeft--
		}
	}

	infrastructure.DisplayGameState(os.Stdout, g.Guessed, g.AttemptsLeft, g.MaxAttempts)
	fmt.Println(g.EndGame())
}

func (g *Game) IsLetterInWord(letter rune) bool {
	return strings.ContainsRune(g.WordToGuess, letter)
}

func (g *Game) UpdateGuessedWord(letter rune) {
	for i, l := range g.WordToGuess {
		if l == letter {
			g.Guessed[i] = letter
		}
	}
}

func (g *Game) IsWordGuessed() bool {
	for _, letter := range g.Guessed {
		if letter == '~' {
			return false
		}
	}

	return true
}

func (g *Game) EndGame() string {
	if g.IsWordGuessed() {
		return fmt.Sprintln("Вы угадали слово:", g.WordToGuess)
	}

	return fmt.Sprintln("Вы проиграли! Загаданное слово было:", g.WordToGuess)
}
