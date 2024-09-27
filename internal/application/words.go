package application

import (
	"crypto/rand"
	"math/big"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
)

func GetRandomWord(category, difficulty, path string) (string, error) {
	// Dictionary contain all words sorted by category and difficulty
	dictionary, err := utils.ReadAndParseWords(path)
	if err != nil {
		return "", err
	}

	words := dictionary.GetVariety(category, difficulty)
	randNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	// Linter made me to do it
	if err != nil {
		return "", err
	}

	return words[randNum.Int64()], nil
}
