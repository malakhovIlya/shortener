package storage

import (
	"context"
	"errors"
)

type Storage interface {
	Save(code, longURL string) error
	Get(code string) (string, error)
}

type InMemoryStorage struct {
	Data map[string]string
}

func (storage InMemoryStorage) Save(code, longURL string) error {
	storage.Data[code] = longURL
	return nil
}

func (storage InMemoryStorage) Get(code string) (string, error) {
	value, ok := storage.Data[code]
	if ok {
		return value, nil
	} else {
		return "", errors.New("URL shortener code not found")
	}
}

func (s *PostgresStorage) Save(code, longURL string) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO links(code, url) VALUES($1, $2)", code, longURL)
	return err
}

func (s *PostgresStorage) Get(code string) (string, error) {
	var url string
	err := s.db.QueryRow(context.Background(),
		"SELECT url FROM links WHERE code=$1", code).Scan(&url)
	return url, err
}
