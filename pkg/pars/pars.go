package pars

import (
	"log"
	"os"
	"strings"
)

func FileToStr(filePath string) []string {
	content := ReadFile(filePath)
	listStrings := strings.Split(string(content), "\n")
	return listStrings
}

func ReadFile(filePath string) []byte {
	log.Printf("Извлечение данных из файла `%v`", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка получения данных:\n%v", err)
	}
	return content
}
