package auth

import (
	"crypto/sha256"
)

func hashing(plain_text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(plain_text))
	return string(hasher.Sum(nil))
}
