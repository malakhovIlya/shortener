package service

import (
	"errors"
	"github.com/google/uuid"
)

type URLShortener struct {
	Storage map[string]string
}

func (urlShortener URLShortener) Shorten(longURL string) string {
	var shortURL = uuid.New().String()[:6]
	urlShortener.Storage[shortURL] = longURL
	return shortURL
}

func (urlShortener URLShortener) Resolve(code string) (string, error) {
	value, ok := urlShortener.Storage[code]
	if ok {
		return value, nil
	} else {
		return "", errors.New("URL shortener code not found")
	}
}
