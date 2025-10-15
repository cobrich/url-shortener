package shortener

import (
	"crypto/rand"
	"io"
)

const customCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateSecureString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	result := make([]byte, length)

	randomBytes := make([]byte, length)

	_, err := io.ReadFull(rand.Reader, randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		randomIndex := int(randomBytes[i]) % len(customCharset)
		result[i] = customCharset[randomIndex]
	}

	return string(result), nil
}
