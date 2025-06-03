package security

import (
	"crypto/rand"
	"fmt"
)

const CharsetAlphabetic = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Charset = CharsetAlphabetic + "0123456789_-"

func GenerateRandomAlphabeticString(n int) (string, error) {
	return GenerateRandomString(n, Charset)
}

func GenerateRandomString(n int, charset string) (string, error) {
	b := make([]byte, n)
	max := byte(len(charset))
	if max == 0 {
		return "", fmt.Errorf("empty charset")
	}

	// Securely fill slice with random indexes
	for i := range b {
		num := make([]byte, 1)
		_, err := rand.Read(num)
		if err != nil {
			return "", err
		}
		b[i] = charset[int(num[0])%len(charset)]
	}

	// Map to actual characters
	result := make([]byte, n)
	for i, v := range b {
		result[i] = charset[v%max]
	}
	return string(result), nil
}
