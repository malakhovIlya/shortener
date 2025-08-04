package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/malakhovIlya/shortener/internal/storage"
)

type URLShortener struct {
	Storage storage.Storage
}

func (urlShortener URLShortener) Shorten(longURL string) string {
	var code = uuid.New().String()[:6]
	err := urlShortener.Storage.Save(code, longURL)
	if err != nil {
		return ""
	}
	return code
}

func (urlShortener URLShortener) Resolve(code string) (string, error) {
	value, err := urlShortener.Storage.Get(code)
	if err != nil {
		return "", errors.New("URL shortener code not found")
	} else {
		return value, nil
	}
}
