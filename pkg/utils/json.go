package utils

import (
	"encoding/json"
	"io"
	"os"
)

type Category struct {
	Category map[string]DifficultyCategory `json:"category"`
}

type DifficultyCategory struct {
	Difficulty map[string][]string `json:"difficulty"`
}

func (data Category) GetVariety(category, difficulty string) []string {
	return data.Category[category].Difficulty[difficulty]
}

// Why this function is so long?
// 'Cause i decided to union opening file, reading file and parsing.
func ReadAndParseWords(path string) (Category, error) {
	// Define a variable of type Category
	var words Category
	// JSON string to unmarshal

	jsonFile, err := os.Open(path)
	if err != nil {
		return words, &FileOpenError{Err: err}
	}

	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal the JSON data into the Go struct
	err = json.Unmarshal(byteValue, &words)
	if err != nil {
		// ("Error unmarshalling JSON:", jsonError)
		return words, &JSONUnmarshalError{Err: err}
	}
	// Print the result
	return words, nil
}
