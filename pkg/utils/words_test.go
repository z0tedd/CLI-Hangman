package utils_test

import (
	"os"
	"slices"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
)

func TestGetRandomWord(t *testing.T) {
	TestJSON := []byte(`
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
        "medium": ["tank", "plane", "missile"],
        "hard": ["robot", "osint", "nuclear"]
      }
    }
  }
}`)
	err := os.WriteFile("test.json", TestJSON, 0o600)
	check(err, t)

	defer os.Remove("test.json")

	var Words utils.Category
	Words, err = utils.ReadAndParseWords("test.json")

	check(err, t)

	variety := Words.GetVariety("fruits", "easy")
	word, err := utils.GetRandomWord("fruits", "easy", "test.json")
	check(err, t)

	if !(slices.Contains(variety, word)) {
		t.Errorf("got %q, expected - %q", variety, word)
	}

	word, err = utils.GetRandomWord("fruits", "easy", "")
	if word != "" {
		t.Error("GetRandomWord doesn't return error")
	}

	if err == nil {
		t.Error("GetRandomWord doesn't return error")
	}
}
