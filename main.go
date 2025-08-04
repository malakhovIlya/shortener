package main

import (
	"github.com/malakhovIlya/shortener/internal/handler"
	"github.com/malakhovIlya/shortener/internal/service"
	"github.com/malakhovIlya/shortener/internal/storage"
	"log"
	"net/http"
)

func main() {
	var localStorage storage.Storage
	connStr := "postgres://user:pass@localhost:5432/shortener?sslmode=disable"
	postgresStorage, errPostgresStorage := storage.NewPostgresStorage(connStr)
	if errPostgresStorage != nil {
		log.Println("Fallback to in-memory:", errPostgresStorage)
		localStorage = storage.InMemoryStorage{Data: make(map[string]string)}
	} else {
		localStorage = postgresStorage
	}

	shortener := &service.URLShortener{Storage: localStorage}

	// Создание обработчика
	h := &handler.Handler{Shortener: shortener}

	// Запуск HTTP сервера
	http.Handle("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
