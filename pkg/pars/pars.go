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
		panic(err)
	}
	listStrings := strings.Split(string(content), "\n")
	return listStrings
}
