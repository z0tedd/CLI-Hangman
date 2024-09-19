package infrastructure

import (
	"fmt"
	"io"
)

var frames = []string{
	`
     _______
     |    \|
           |
           |
           |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
           |
           |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
     |     |
           |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
    /|     |
           |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
    /|\    |
           |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
    /|\    |
    /      |
          /|
    =========`,
	`
     _______
     |    \|
     O     |
    /|\    |
    / \    |
          /|
    =========`,
}

func DisplayGameState(w io.Writer, guessed []rune, attemptsLeft, maxAttempts int) {
	amountofFrames := len(frames)

	fmt.Fprintln(w, "\nСлово: ", string(guessed))

	stageIndex := maxAttempts - attemptsLeft
	gap := (float64(maxAttempts)) / float64((amountofFrames-1)-1)
	// amountofFrames-2 is used, because i print last frame in the end of the game
	// and don't want to draw it when attemptsLeft = 1
	// main idea was to draw it in the last moment
	defer fmt.Fprintf(w, "Осталось попыток: %d\n\n", attemptsLeft)

	if attemptsLeft == 0 {
		fmt.Fprintln(w, frames[len(frames)-1])
		return
	}

	for i := range amountofFrames {
		if float64(stageIndex) <= gap*float64(i) {
			fmt.Fprintln(w, frames[i])
			return
		}
	}
}
