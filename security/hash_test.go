package security

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashBcryptAndCheck(t *testing.T) {
	password := "my-secret-password"

	// Hash generieren
	hash, err := HashBcrypt(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("unexpected error while hashing: %v", err)
	}

	// Check sollte true zurückgeben
	if !CheckBcryptHash(password, hash) {
		t.Error("CheckBcryptHash returned false for correct password")
	}

	// Check mit falschem Passwort
	if CheckBcryptHash("wrong-password", hash) {
		t.Error("CheckBcryptHash returned true for incorrect password")
	}
}

func TestHashBcrypt_InvalidCost(t *testing.T) {
	_, err := HashBcrypt("test", 1_000_000) // ungültiger/extremer Cost
	if err == nil {
		t.Error("expected error with invalid cost, got nil")
	}
}
