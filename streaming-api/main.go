package main

import (
	"log"

	"github.com/MishaAnikutin/streaming-api/handlers"
	"github.com/MishaAnikutin/streaming-api/server"
)

func main() {
	port := "80" // TODO: брать из окружения

	server := server.New(port, handlers.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s", err.Error())
	}

	defer server.Stop()
}
