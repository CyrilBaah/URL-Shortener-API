package storage

import (
	"math/rand"
	"sync"
	"time"
)

var store = make(map[string]string)
var mu sync.Mutex

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	short := make([]byte, 6)
	for i := range short {
		short[i] = letters[rand.Intn(len(letters))]
	}
	return string(short)
}

func SaveURL(originalURL string) string {
	mu.Lock()
	defer mu.Unlock()

	shortURL := generateShortURL()
	store[shortURL] = originalURL
	return shortURL
}

func GetOriginalURL(shortURL string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	originalURL, exists := store[shortURL]
	return originalURL, exists
}
