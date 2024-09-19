package utils

import (
	"math/rand"
)

func GetRandomWord(category string, difficulty string, path string) (string, error) {
	// Dictionary contain all words sorted by category and difficulty
	dictionary, err := ReadAndParseWords(path)
	// fmt.Println(words, dictionary, category, difficulty)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	words := dictionary.GetVariety(category, difficulty)
	return words[rand.Intn(len(words))], err
}
