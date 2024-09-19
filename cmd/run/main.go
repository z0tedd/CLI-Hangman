package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend-academy_2024_project_1-go-z0tedd/pkg/utils"
)

func main() {
	// Инициализация
	// rand.Seed(time.Now().UnixNano())  - deprecated
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	fmt.Println("Выведите путь до файла:")
	var path string
	if path == "" {
		path = "words.json"
	}
	fmt.Scanln(&path)
	fmt.Println("Выведите название категории: fruits|weapons для стандартного конфига")
	var category string
	fmt.Scanln(&category)
	// Выбор уровня сложности
	var difficulty string
	fmt.Println("Выберите уровень сложности: easy, medium, hard")
	fmt.Scanln(&difficulty)
	Hardness := map[string]int{"easy": 1, "medium": 2, "hard": 3}
	// Получение случайного слова
	word, err := utils.GetRandomWord(category, difficulty, path)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	// Запуск игры
	game := application.NewGame(word, Hardness[difficulty])
	game.Start()
}
