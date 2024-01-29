package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func CreateShortURL(baseUrl, longURL string) string {
	hash := sha1.New()
	hash.Write([]byte(longURL))
	hashBytes := hash.Sum(nil)
	// Encode the hash into base64 to create a short key
	hashString := base64.URLEncoding.EncodeToString(hashBytes)

	shortKey := hashString[:14]
	// Use the short key as a path component in the short URL
	shortURL := baseUrl + "/" + shortKey
	return shortURL
}
