package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomBase64Hash(size int) (string, error) {
	// Create a byte slice of the specified size
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// Convert to Base64 string
	return base64.URLEncoding.EncodeToString(bytes), nil
}
