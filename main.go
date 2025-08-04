package main

import (
	"github.com/malakhovIlya/shortener/internal/handler"
	"github.com/malakhovIlya/shortener/internal/service"
	"github.com/malakhovIlya/shortener/internal/storage"
	"net/http"
)

func main() {
	memoryStorage := storage.InMemoryStorage{Data: make(map[string]string)}
	shortener := &service.URLShortener{Storage: memoryStorage.Data}

	// Создание обработчика
	h := &handler.Handler{Shortener: shortener}

	// Запуск HTTP сервера
	http.Handle("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
