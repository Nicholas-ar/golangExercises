package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
)

// URLShortener handles shortening and expanding URLs
type URLShortener struct {
	urls  map[string]string
	mutex sync.RWMutex
}

// NewURLShortener creates a new URLShortener instance
func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}

// generateShortCode creates a random short code
func generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

var shortener = NewURLShortener()

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortener.mutex.Lock()
	shortCode := generateShortCode()
	shortener.urls[shortCode] = url
	shortener.mutex.Unlock()

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortCode)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	if shortCode == "" {
		fmt.Fprint(w, "Welcome to URL Shortener! Use GET /shorten?url=<your-url> to create short URLs")
		return
	}

	shortener.mutex.RLock()
	originalURL, exists := shortener.urls[shortCode]
	shortener.mutex.RUnlock()

	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("URL Shortener starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
