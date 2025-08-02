package main

import (
	"github.com/malakhovIlya/shortener/internal/handler"
	"github.com/malakhovIlya/shortener/internal/service"
	"net/http"
)

func main() {
	// Создание сервиса сокращения URL
	shortener := &service.URLShortener{Storage: make(map[string]string)}

	// Создание обработчика
	h := &handler.Handler{Shortener: shortener}

	// Запуск HTTP сервера
	http.Handle("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
