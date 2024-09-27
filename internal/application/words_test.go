package application_test

import (
	"os"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err, err)

	defer os.Remove("test.json")

	Words, err := utils.ReadAndParseWords("test.json")

	assert.Nil(t, err, err)

	variety := Words.GetVariety("fruits", "easy")

	word, err := application.GetRandomWord("fruits", "easy", "test.json")
	assert.Nil(t, err, err)

	assert.Contains(t, variety, word, "got %q, expected - %q", variety, word)

	word, err = application.GetRandomWord("fruits", "easy", "")
	assert.Equal(t, word, "", "GetRandomWord must return error in this example")

	assert.NotNil(t, err, "GetRandomWord must return error in this example")
}
