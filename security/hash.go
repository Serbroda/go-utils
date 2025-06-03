package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashBcrypt(plain string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	return string(bytes), err
}

func CheckBcryptHash(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
