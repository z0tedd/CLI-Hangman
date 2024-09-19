package infrastructure

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDisplayGameState(t *testing.T) {
	var b bytes.Buffer
	guessed := []rune("~~~~~~")
	attemtsleft := 12
	maxattempts := 12
	DisplayGameState(&b, guessed, attemtsleft, maxattempts)
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

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	b.Reset()
	DisplayGameState(&b, guessed, 0, maxattempts)
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
	if got != want {
		fmt.Println(got)
		fmt.Println(want)
		t.Errorf("got = %q, want %q", got, want)
	}
}
