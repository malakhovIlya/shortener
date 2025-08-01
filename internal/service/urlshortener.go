package service

import (
	"errors"
	"github.com/google/uuid"
)

type URLShortener struct {
	storage map[string]string
}

func (urlShortener URLShortener) Shorten(longURL string) string {
	var shortURL = uuid.New().String()[:6]
	urlShortener.storage[shortURL] = longURL
	return shortURL
}

func (urlShortener URLShortener) Resolve(code string) (string, error) {
	value, ok := urlShortener.storage[code]
	if ok {
		return value, nil
	} else {
		return "", errors.New("URL shortener code not found")
	}
}
