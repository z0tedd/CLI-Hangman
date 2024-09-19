package utils

import (
	"os"
	"slices"
	"testing"
)

func TestGetRandomWord(t *testing.T) {
	TestJson := []byte(`
{
  "category": {
    "fruits": {
      "difficulty": {
        "easy": ["apple", "banana", "melon"],
        "medium": ["cherry", "strawberry", "lemon"],
        "hard": ["dragonfruit", "cactus", "mango"]
      }
    },
    "weapons": {
      "difficulty": {
        "easy": ["knife", "sword", "gun"],
        "medium": ["tank", "plane", "missle"],
        "hard": ["robot", "osint", "nuclear"]
      }
    }
  }
}`)
	err := os.WriteFile("test.json", TestJson, 0644)
	check(err, t)
	defer os.Remove("test.json")
	var Words Category
	Words, err = ReadAndParseWords("test.json")

	check(err, t)
	variety := Words.GetVariety("fruits", "easy")
	word, err := GetRandomWord("fruits", "easy", "test.json")
	check(err, t)
	if !(slices.Contains(variety, word)) {
		t.Errorf("got %q, expected - %q", variety, word)
	}
	word, err = GetRandomWord("fruits", "easy", "")
	if word != "" {
		t.Error("GetRandomWord doesn't return error")
	}
	if err == nil {
		t.Error("GetRandomWord doesn't return error")
	}
}
