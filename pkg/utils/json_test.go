package utils_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/domain"
	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
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
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove("test.json")

	var Words domain.Category

	Words, err = utils.ReadAndParseWords("test.json")
	if err != nil {
		t.Fatal(err)
	}

	got := Words.GetVariety("fruits", "easy")

	want := []string{"apple", "banana", "melon"}
	if !(reflect.DeepEqual(got, want)) {
		t.Fatalf("got: %q, want %q", got, want)
	}

	got = Words.GetVariety("weapons", "easy")

	want = []string{"knife", "sword", "gun"}

	if !(reflect.DeepEqual(got, want)) {
		t.Fatalf("got: %q, want %q", got, want)
	}
}

func TestReadAndParseWords(t *testing.T) {
	err := os.WriteFile("test.json", TestJSON, 0o600)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove("test.json")

	Words, err := utils.ReadAndParseWords("test.json")
	if err != nil {
		t.Fatal(err)
	}

	check := Words.Category["fruits"].Difficulty["easy"][0]
	if check != "apple" {
		t.Errorf("got %q, expected - %q", check, "apple")
	}
}

func TestReadAndParseWordsJsonError(t *testing.T) {
	ErrorJSON := "{name: test}" // random json data, unmarshall must return error

	err := os.WriteFile("test.json", []byte(ErrorJSON), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove("test.json")

	var Words domain.Category

	Words, err = utils.ReadAndParseWords("test.json")
	if err == nil {
		t.Fatalf("function doesn't return error, words - %q, err - %q", Words, err)
	}
}
