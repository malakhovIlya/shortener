package service

import (
	"github.com/malakhovIlya/shortener/internal/storage"
	"testing"
)

func TestURLShortener_Shorten(t *testing.T) {
	shortener := &URLShortener{
		Storage: storage.InMemoryStorage{
			Data: make(map[string]string),
		},
	}

	url := "https://example.com"
	code := shortener.Shorten(url)

	if len(code) != 6 {
		t.Errorf("Expected code length 6, got %d", len(code))
	}

	code2 := shortener.Shorten(url)
	if code == code2 {
		t.Error("Expected different codes for repeated calls")
	}
}

func TestURLShortener_Resolve(t *testing.T) {
	shortener := &URLShortener{
		Storage: storage.InMemoryStorage{
			Data: map[string]string{
				"code123": "https://example.com",
			},
		},
	}

	expectedUrl := "https://example.com"
	code := "code123"
	url, err := shortener.Resolve(code)

	if err != nil {
		t.Errorf("Une erreur s'est produite, %s", err)
	}

	if url != expectedUrl {
		t.Errorf("Expected url - %s", expectedUrl)
	}
}

func TestURLShortener_Resolve_NotFound(t *testing.T) {
	shortener := &URLShortener{
		Storage: storage.InMemoryStorage{
			Data: map[string]string{
				"code123": "https://example.com",
			},
		},
	}
	code := "code12"
	_, err := shortener.Resolve(code)
	if err == nil {
		t.Errorf("Erreur attendue")
	}
}
