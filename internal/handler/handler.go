package handler

import (
	"encoding/json"
	"github.com/malakhovIlya/shortener/internal/service"
	"net/http"
	"strings"
)

type Handler struct {
	Shortener *service.URLShortener
}

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Code string `json:"code"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/shorten" {
		var req Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		code := h.Shortener.Shorten(req.URL)
		resp := Response{Code: code}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/r/") {
		code := strings.TrimPrefix(r.URL.Path, "/r/")
		if code == "" {
			http.Error(w, "Code is required", http.StatusBadRequest)
			return
		}

		originalURL, err := h.Shortener.Resolve(code)
		if err != nil {
			http.Error(w, "Error resolving URL: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if originalURL == "" {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusFound)
	}
	http.Error(w, "Not found", http.StatusNotFound)
}
