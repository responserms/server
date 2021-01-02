package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateEncryptionKey() (string, error) {
	rawKey := make([]byte, 32)

	_, err := rand.Read(rawKey)
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(rawKey), nil
}
