package utils

import (
	"os"
	"reflect"
	"testing"
)

var TestJson = []byte(`
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

func check(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetVariety(t *testing.T) {
	err := os.WriteFile("test.json", TestJson, 0644)
	check(err, t)
	defer os.Remove("test.json")
	var Words Category
	Words, err = ReadAndParseWords("test.json")
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
	err := os.WriteFile("test.json", TestJson, 0644)
	check(err, t)
	defer os.Remove("test.json")
	var Words Category
	Words, err = ReadAndParseWords("test.json")

	check(err, t)
	check := Words.Category["fruits"].Difficulty["easy"][0]
	if check != "apple" {
		t.Errorf("got %q, expected - %q", check, "apple")
	}
}

func TestReadAndParseWordsJsonError(t *testing.T) {
	ErrorJson := "{name: test}" // random json data, unmarshall must return error
	err := os.WriteFile("test.json", []byte(ErrorJson), 0644)
	check(err, t)
	defer os.Remove("test.json")
	var Words Category
	Words, err = ReadAndParseWords("test.json")
	if err == nil {
		t.Fatalf("function doesn't return error, words - %q, err - %q", Words, err)
	}
}
