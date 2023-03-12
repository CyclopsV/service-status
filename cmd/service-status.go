package main

import (
	"github.com/CyclopsV/service-status-skillbox/api/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Запуск сервера")
	r := routes.CreateRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
