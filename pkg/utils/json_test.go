package utils_test

import (
	"os"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var TestJSON = []byte(`
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

func TestGetVariety(t *testing.T) {
	err := os.WriteFile("test.json", TestJSON, 0o600)
	assert.Nil(t, err, err)

	defer os.Remove("test.json")

	Words, err := utils.ReadAndParseWords("test.json")

	assert.Nil(t, err, err)

	got := Words.GetVariety("fruits", "easy")

	want := []string{"apple", "banana", "melon"}
	assert.Equal(t, got, want, "got: %q, want %q", got, want)

	got = Words.GetVariety("weapons", "easy")

	want = []string{"knife", "sword", "gun"}

	assert.Equal(t, got, want, "got: %q, want %q", got, want)
}

func TestReadAndParseWords(t *testing.T) {
	err := os.WriteFile("test.json", TestJSON, 0o600)
	assert.Nil(t, err, err)

	defer os.Remove("test.json")

	Words, err := utils.ReadAndParseWords("test.json")
	assert.Nil(t, err, err)

	check := Words.Category["fruits"].Difficulty["easy"][0]
	assert.Equal(t, check, "apple", "got %q, expected - %q", check, "apple")
}

func TestReadAndParseWordsJsonError(t *testing.T) {
	ErrorJSON := "{name: test}" // random json data, unmarshall must return error

	err := os.WriteFile("test.json", []byte(ErrorJSON), 0o600)
	assert.Nil(t, err, err)

	defer os.Remove("test.json")

	Words, err := utils.ReadAndParseWords("test.json")

	assert.NotNil(t, err, "function doesn't return error, words - %q, err - %q", Words, err)
}
