package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func CreateShortURL(longURL string) string {
	hash := sha1.New()
	hash.Write([]byte(longURL))
	hashBytes := hash.Sum(nil)
	// Encode the hash into base64 to create a short key
	hashString := base64.URLEncoding.EncodeToString(hashBytes)

	shortURL := hashString[:8]
	// Use the short key as a path component in the short URL
	return shortURL
}
