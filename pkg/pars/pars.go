package pars

import (
	"log"
	"os"
	"strings"
)

func ParseFile(filePath string) []string {
	log.Printf("Извлечение данных из файла `%v`", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка получения данных SMS:\n%v", err)
	}
	listStrings := strings.Split(string(content), "\n")
	return listStrings
}
